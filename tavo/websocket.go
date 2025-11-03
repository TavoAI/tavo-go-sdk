package tavo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocketConfig holds configuration for WebSocket connections
type WebSocketConfig struct {
	// ReconnectInterval is the time to wait before attempting to reconnect
	ReconnectInterval time.Duration
	// MaxReconnectAttempts is the maximum number of reconnection attempts
	MaxReconnectAttempts int
	// PingInterval is the interval for sending ping messages
	PingInterval time.Duration
	// ReadTimeout is the read timeout for WebSocket messages
	ReadTimeout time.Duration
	// WriteTimeout is the write timeout for WebSocket messages
	WriteTimeout time.Duration
}

// DefaultWebSocketConfig returns a default WebSocket configuration
func DefaultWebSocketConfig() WebSocketConfig {
	return WebSocketConfig{
		ReconnectInterval:    5 * time.Second,
		MaxReconnectAttempts: 10,
		PingInterval:         30 * time.Second,
		ReadTimeout:          60 * time.Second,
		WriteTimeout:         10 * time.Second,
	}
}

// WebSocketConnection represents a WebSocket connection with automatic reconnection
type WebSocketConnection struct {
	config        WebSocketConfig
	baseURL       string
	apiKey        string
	jwtToken      string
	sessionToken  string
	conn          *websocket.Conn
	url           string
	isConnected   bool
	mu            sync.RWMutex
	ctx           context.Context
	cancel        context.CancelFunc
	reconnectChan chan bool
}

// NewWebSocketConnection creates a new WebSocket connection
func NewWebSocketConnection(baseURL, apiKey string, config WebSocketConfig) *WebSocketConnection {
	ctx, cancel := context.WithCancel(context.Background())
	return &WebSocketConnection{
		config:        config,
		baseURL:       baseURL,
		apiKey:        apiKey,
		isConnected:   false,
		ctx:           ctx,
		cancel:        cancel,
		reconnectChan: make(chan bool, 1),
	}
}

// SetJWTToken sets the JWT token for authentication
func (ws *WebSocketConnection) SetJWTToken(token string) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.jwtToken = token
}

// SetSessionToken sets the session token for authentication
func (ws *WebSocketConnection) SetSessionToken(token string) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.sessionToken = token
}

// Connect establishes the WebSocket connection
func (ws *WebSocketConnection) Connect() error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if ws.isConnected {
		return fmt.Errorf("already connected")
	}

	// Parse base URL and convert to WebSocket URL
	u, err := url.Parse(ws.baseURL)
	if err != nil {
		return fmt.Errorf("invalid base URL: %w", err)
	}

	// Convert HTTP/HTTPS to WS/WSS
	if u.Scheme == "https" {
		u.Scheme = "wss"
	} else {
		u.Scheme = "ws"
	}

	ws.url = u.String()

	// Create WebSocket dialer
	dialer := websocket.DefaultDialer
	dialer.HandshakeTimeout = 10 * time.Second

	// Add authentication headers
	header := make(map[string][]string)
	if ws.jwtToken != "" {
		header["Authorization"] = []string{fmt.Sprintf("Bearer %s", ws.jwtToken)}
	} else if ws.sessionToken != "" {
		header["X-Session-Token"] = []string{ws.sessionToken}
	} else if ws.apiKey != "" {
		header["X-API-Key"] = []string{ws.apiKey}
	}

	// Establish connection
	conn, _, err := dialer.Dial(ws.url, header)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	ws.conn = conn
	ws.isConnected = true

	// Start ping/pong handler
	go ws.pingHandler()

	return nil
}

// Disconnect closes the WebSocket connection
func (ws *WebSocketConnection) Disconnect() error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if !ws.isConnected {
		return nil
	}

	ws.isConnected = false
	ws.cancel() // Cancel context to stop all goroutines

	if ws.conn != nil {
		err := ws.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Printf("Error sending close message: %v", err)
		}
		ws.conn.Close()
		ws.conn = nil
	}

	return nil
}

// IsConnected returns whether the connection is currently active
func (ws *WebSocketConnection) IsConnected() bool {
	ws.mu.RLock()
	defer ws.mu.RUnlock()
	return ws.isConnected
}

// pingHandler maintains the connection with periodic ping messages
func (ws *WebSocketConnection) pingHandler() {
	ticker := time.NewTicker(ws.config.PingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ws.ctx.Done():
			return
		case <-ticker.C:
			ws.mu.RLock()
			if !ws.isConnected || ws.conn == nil {
				ws.mu.RUnlock()
				return
			}

			ws.conn.SetWriteDeadline(time.Now().Add(ws.config.WriteTimeout))
			if err := ws.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("Failed to send ping: %v", err)
				ws.mu.RUnlock()
				go ws.handleReconnect()
				return
			}
			ws.mu.RUnlock()
		}
	}
}

// handleReconnect attempts to reconnect to the WebSocket
func (ws *WebSocketConnection) handleReconnect() {
	ws.mu.Lock()
	if !ws.isConnected {
		ws.mu.Unlock()
		return
	}
	ws.isConnected = false
	if ws.conn != nil {
		ws.conn.Close()
		ws.conn = nil
	}
	ws.mu.Unlock()

	for attempt := 1; attempt <= ws.config.MaxReconnectAttempts; attempt++ {
		select {
		case <-ws.ctx.Done():
			return
		case <-time.After(ws.config.ReconnectInterval):
			log.Printf("Attempting to reconnect (attempt %d/%d)", attempt, ws.config.MaxReconnectAttempts)

			if err := ws.Connect(); err != nil {
				log.Printf("Reconnection attempt %d failed: %v", attempt, err)
				continue
			}

			log.Printf("Successfully reconnected on attempt %d", attempt)
			return
		}
	}

	log.Printf("Failed to reconnect after %d attempts", ws.config.MaxReconnectAttempts)
}

// SendMessage sends a message over the WebSocket connection
func (ws *WebSocketConnection) SendMessage(messageType string, data interface{}) error {
	ws.mu.RLock()
	defer ws.mu.RUnlock()

	if !ws.isConnected || ws.conn == nil {
		return fmt.Errorf("not connected")
	}

	message := map[string]interface{}{
		"type":      messageType,
		"data":      data,
		"timestamp": time.Now().Unix(),
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	ws.conn.SetWriteDeadline(time.Now().Add(ws.config.WriteTimeout))
	return ws.conn.WriteMessage(websocket.TextMessage, jsonData)
}

// ReadMessage reads a message from the WebSocket connection
func (ws *WebSocketConnection) ReadMessage() ([]byte, error) {
	ws.mu.RLock()
	conn := ws.conn
	ws.mu.RUnlock()

	if conn == nil {
		return nil, fmt.Errorf("not connected")
	}

	conn.SetReadDeadline(time.Now().Add(ws.config.ReadTimeout))
	_, message, err := conn.ReadMessage()
	if err != nil {
		// If read fails, trigger reconnection
		go ws.handleReconnect()
		return nil, fmt.Errorf("failed to read message: %w", err)
	}

	return message, nil
}

// ReadMessages reads messages in a loop and sends them to a channel
func (ws *WebSocketConnection) ReadMessages(messageChan chan<- []byte, errorChan chan<- error) {
	for {
		select {
		case <-ws.ctx.Done():
			close(messageChan)
			return
		default:
			message, err := ws.ReadMessage()
			if err != nil {
				select {
				case errorChan <- err:
				case <-time.After(time.Second):
					// Don't block if error channel is full
				}
				return
			}

			select {
			case messageChan <- message:
			case <-time.After(time.Second):
				// Don't block if message channel is full
			}
		}
	}
}

// WebSocketOperations provides high-level WebSocket operations for the Tavo client
type WebSocketOperations struct {
	client *Client
}

// NewWebSocketOperations creates a new WebSocket operations instance
func NewWebSocketOperations(client *Client) *WebSocketOperations {
	return &WebSocketOperations{client: client}
}

// ConnectToScanProgress connects to real-time scan progress updates for a specific scan
func (ws *WebSocketOperations) ConnectToScanProgress(scanID string, config WebSocketConfig) (*WebSocketConnection, error) {
	baseURL := ws.client.GetBaseURL()
	clientConfig := ws.client.GetConfig()

	conn := NewWebSocketConnection(baseURL, clientConfig.APIKey, config)
	conn.SetJWTToken(clientConfig.JWTToken)
	conn.SetSessionToken(clientConfig.SessionToken)

	// Set the specific endpoint for scan progress
	conn.url = fmt.Sprintf("%s/api/v1/code/scans/%s/progress", baseURL, scanID)
	if baseURL[:5] == "https" {
		conn.url = "wss" + conn.url[5:]
	} else {
		conn.url = "ws" + conn.url[4:]
	}

	if err := conn.Connect(); err != nil {
		return nil, fmt.Errorf("failed to connect to scan progress: %w", err)
	}

	return conn, nil
}

// ConnectToGeneralUpdates connects to general real-time updates
func (ws *WebSocketOperations) ConnectToGeneralUpdates(config WebSocketConfig) (*WebSocketConnection, error) {
	baseURL := ws.client.GetBaseURL()
	clientConfig := ws.client.GetConfig()

	conn := NewWebSocketConnection(baseURL, clientConfig.APIKey, config)
	conn.SetJWTToken(clientConfig.JWTToken)
	conn.SetSessionToken(clientConfig.SessionToken)

	if err := conn.Connect(); err != nil {
		return nil, fmt.Errorf("failed to connect to general updates: %w", err)
	}

	return conn, nil
}
