package tavo

import "fmt"

// WebhookOperations handles webhook operations
// Deprecated: Use GitHub App webhook management instead for CLI tools
type WebhookOperations struct {
	client *Client
}

// ListWebhooks lists all webhooks
// Deprecated: Use GitHub App webhook management instead
func (w *WebhookOperations) ListWebhooks(params map[string]interface{}) (map[string]interface{}, error) {
	query := ""
	if params != nil {
		query = "?"
		for key, value := range params {
			if query != "?" {
				query += "&"
			}
			query += fmt.Sprintf("%s=%v", key, value)
		}
	}
	return w.client.makeRequest("GET", "/webhooks"+query, nil)
}

// GetWebhook returns a specific webhook
// Deprecated: Use GitHub App webhook management instead
func (w *WebhookOperations) GetWebhook(webhookID string) (map[string]interface{}, error) {
	return w.client.makeRequest("GET", "/webhooks/"+webhookID, nil)
}

// CreateWebhook creates a new webhook
// Deprecated: Use GitHub App webhook management instead
func (w *WebhookOperations) CreateWebhook(webhookData map[string]interface{}) (map[string]interface{}, error) {
	return w.client.makeRequest("POST", "/webhooks", webhookData)
}

// UpdateWebhook updates a webhook
// Deprecated: Use GitHub App webhook management instead
func (w *WebhookOperations) UpdateWebhook(webhookID string, webhookData map[string]interface{}) (map[string]interface{}, error) {
	return w.client.makeRequest("PUT", "/webhooks/"+webhookID, webhookData)
}

// DeleteWebhook deletes a webhook
// Deprecated: Use GitHub App webhook management instead
func (w *WebhookOperations) DeleteWebhook(webhookID string) error {
	_, err := w.client.makeRequest("DELETE", "/webhooks/"+webhookID, nil)
	return err
}

// TestWebhook tests a webhook by sending a test payload
// Deprecated: Use GitHub App webhook management instead
func (w *WebhookOperations) TestWebhook(webhookID string) (map[string]interface{}, error) {
	return w.client.makeRequest("POST", "/webhooks/"+webhookID+"/test", nil)
}

// GetWebhookDeliveries returns delivery history for a webhook
// Deprecated: Use GitHub App webhook management instead
func (w *WebhookOperations) GetWebhookDeliveries(webhookID string, params map[string]interface{}) (map[string]interface{}, error) {
	query := ""
	if params != nil {
		query = "?"
		for key, value := range params {
			if query != "?" {
				query += "&"
			}
			query += fmt.Sprintf("%s=%v", key, value)
		}
	}
	return w.client.makeRequest("GET", "/webhooks/"+webhookID+"/deliveries"+query, nil)
}
