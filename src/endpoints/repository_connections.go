package tavo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// RepositoryConnectionsClient handles repository_connections API calls
type RepositoryConnectionsClient struct {
	client *Client
}

// PostRoot POST /
func (c *RepositoryConnectionsClient) PostRoot(connection_in interface{}) (interface{}, error) {
		url := fmt.Sprintf("/", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyBytes, err := json.Marshal(connection_in)
		if err != nil {
			return nil, err
		}
		body := bytes.NewReader(bodyBytes)
		req, err := http.NewRequest("POST", fullURL, body)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "tavo-sdk-go/0.1.0")
		if c.client.apiKey != "" {
			req.Header.Set("X-API-Key", c.client.apiKey)
		} else if c.client.deviceToken != "" {
			req.Header.Set("Authorization", "Bearer "+c.client.deviceToken)
		}
		resp, err := c.client.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
		}
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var result interface{}
		err = json.Unmarshal(bodyBytes, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
}
// GetRoot GET /
func (c *RepositoryConnectionsClient) GetRoot(provider_id *string, connection_type *string, is_active *bool) (interface{}, error) {
		url := fmt.Sprintf("/", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("provider_id", fmt.Sprintf("%v", provider_id))
		params.Add("connection_type", fmt.Sprintf("%v", connection_type))
		params.Add("is_active", fmt.Sprintf("%v", is_active))
		fullURL += "?" + params.Encode()
		body := (*bytes.Reader)(nil)
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "tavo-sdk-go/0.1.0")
		if c.client.apiKey != "" {
			req.Header.Set("X-API-Key", c.client.apiKey)
		} else if c.client.deviceToken != "" {
			req.Header.Set("Authorization", "Bearer "+c.client.deviceToken)
		}
		resp, err := c.client.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
		}
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var result interface{}
		err = json.Unmarshal(bodyBytes, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
}
// Get{connection_id} GET /{connection_id}
func (c *RepositoryConnectionsClient) Get{connection_id}(connection_id string) (interface{}, error) {
		url := fmt.Sprintf("/{connection_id}", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("connection_id", fmt.Sprintf("%v", connection_id))
		fullURL += "?" + params.Encode()
		body := (*bytes.Reader)(nil)
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "tavo-sdk-go/0.1.0")
		if c.client.apiKey != "" {
			req.Header.Set("X-API-Key", c.client.apiKey)
		} else if c.client.deviceToken != "" {
			req.Header.Set("Authorization", "Bearer "+c.client.deviceToken)
		}
		resp, err := c.client.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
		}
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var result interface{}
		err = json.Unmarshal(bodyBytes, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
}
// Put{connection_id} PUT /{connection_id}
func (c *RepositoryConnectionsClient) Put{connection_id}(connection_id string, connection_update interface{}) (interface{}, error) {
		url := fmt.Sprintf("/{connection_id}", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyData := map[string]interface{}{
			"connection_id": connection_id,
			"connection_update": connection_update,
		}
		bodyBytes, err := json.Marshal(bodyData)
		if err != nil {
			return nil, err
		}
		body := bytes.NewReader(bodyBytes)
		req, err := http.NewRequest("PUT", fullURL, body)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "tavo-sdk-go/0.1.0")
		if c.client.apiKey != "" {
			req.Header.Set("X-API-Key", c.client.apiKey)
		} else if c.client.deviceToken != "" {
			req.Header.Set("Authorization", "Bearer "+c.client.deviceToken)
		}
		resp, err := c.client.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
		}
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var result interface{}
		err = json.Unmarshal(bodyBytes, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
}
// Delete{connection_id} DELETE /{connection_id}
func (c *RepositoryConnectionsClient) Delete{connection_id}(connection_id string) (interface{}, error) {
		url := fmt.Sprintf("/{connection_id}", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyBytes, err := json.Marshal(connection_id)
		if err != nil {
			return nil, err
		}
		body := bytes.NewReader(bodyBytes)
		req, err := http.NewRequest("DELETE", fullURL, body)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "tavo-sdk-go/0.1.0")
		if c.client.apiKey != "" {
			req.Header.Set("X-API-Key", c.client.apiKey)
		} else if c.client.deviceToken != "" {
			req.Header.Set("Authorization", "Bearer "+c.client.deviceToken)
		}
		resp, err := c.client.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
		}
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var result interface{}
		err = json.Unmarshal(bodyBytes, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
}
// Post{connection_id}validate POST /{connection_id}/validate
func (c *RepositoryConnectionsClient) Post{connection_id}validate(connection_id string) (interface{}, error) {
		url := fmt.Sprintf("/{connection_id}/validate", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyBytes, err := json.Marshal(connection_id)
		if err != nil {
			return nil, err
		}
		body := bytes.NewReader(bodyBytes)
		req, err := http.NewRequest("POST", fullURL, body)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "tavo-sdk-go/0.1.0")
		if c.client.apiKey != "" {
			req.Header.Set("X-API-Key", c.client.apiKey)
		} else if c.client.deviceToken != "" {
			req.Header.Set("Authorization", "Bearer "+c.client.deviceToken)
		}
		resp, err := c.client.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
		}
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var result interface{}
		err = json.Unmarshal(bodyBytes, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
}
// Post{connection_id}refresh POST /{connection_id}/refresh
func (c *RepositoryConnectionsClient) Post{connection_id}refresh(connection_id string) (interface{}, error) {
		url := fmt.Sprintf("/{connection_id}/refresh", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyBytes, err := json.Marshal(connection_id)
		if err != nil {
			return nil, err
		}
		body := bytes.NewReader(bodyBytes)
		req, err := http.NewRequest("POST", fullURL, body)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "tavo-sdk-go/0.1.0")
		if c.client.apiKey != "" {
			req.Header.Set("X-API-Key", c.client.apiKey)
		} else if c.client.deviceToken != "" {
			req.Header.Set("Authorization", "Bearer "+c.client.deviceToken)
		}
		resp, err := c.client.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
		}
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var result interface{}
		err = json.Unmarshal(bodyBytes, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
}
