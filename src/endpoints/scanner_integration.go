package tavo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// ScannerIntegrationClient handles scanner_integration API calls
type ScannerIntegrationClient struct {
	client *Client
}

// Getrulesdiscover GET /rules/discover
func (c *ScannerIntegrationClient) Getrulesdiscover(category *string, language *string, scanner_type *string, limit *float64) (interface{}, error) {
		url := fmt.Sprintf("/rules/discover", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("category", fmt.Sprintf("%v", category))
		params.Add("language", fmt.Sprintf("%v", language))
		params.Add("scanner_type", fmt.Sprintf("%v", scanner_type))
		params.Add("limit", fmt.Sprintf("%v", limit))
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
// Getrulesbundle{bundle_id}rules GET /rules/bundle/{bundle_id}/rules
func (c *ScannerIntegrationClient) Getrulesbundle{bundle_id}rules(bundle_id string, severity *string, language *string, limit *float64) (interface{}, error) {
		url := fmt.Sprintf("/rules/bundle/{bundle_id}/rules", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("bundle_id", fmt.Sprintf("%v", bundle_id))
		params.Add("severity", fmt.Sprintf("%v", severity))
		params.Add("language", fmt.Sprintf("%v", language))
		params.Add("limit", fmt.Sprintf("%v", limit))
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
// Postrulesbundle{bundle_id}use POST /rules/bundle/{bundle_id}/use
func (c *ScannerIntegrationClient) Postrulesbundle{bundle_id}use(bundle_id string, scan_id *string) (interface{}, error) {
		url := fmt.Sprintf("/rules/bundle/{bundle_id}/use", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyData := map[string]interface{}{
			"bundle_id": bundle_id,
			"scan_id": scan_id,
		}
		bodyBytes, err := json.Marshal(bodyData)
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
// Getpluginsdiscover GET /plugins/discover
func (c *ScannerIntegrationClient) Getpluginsdiscover(plugin_type *string, language *string, scanner_integration *bool, limit *float64) (interface{}, error) {
		url := fmt.Sprintf("/plugins/discover", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("plugin_type", fmt.Sprintf("%v", plugin_type))
		params.Add("language", fmt.Sprintf("%v", language))
		params.Add("scanner_integration", fmt.Sprintf("%v", scanner_integration))
		params.Add("limit", fmt.Sprintf("%v", limit))
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
// Getplugins{plugin_id}config GET /plugins/{plugin_id}/config
func (c *ScannerIntegrationClient) Getplugins{plugin_id}config(plugin_id string) (interface{}, error) {
		url := fmt.Sprintf("/plugins/{plugin_id}/config", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("plugin_id", fmt.Sprintf("%v", plugin_id))
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
// Postscannerheartbeat POST /scanner/heartbeat
func (c *ScannerIntegrationClient) Postscannerheartbeat(scanner_version string, scanner_type *string, active_rules *[]string, active_plugins *[]string) (interface{}, error) {
		url := fmt.Sprintf("/scanner/heartbeat", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyData := map[string]interface{}{
			"scanner_version": scanner_version,
			"scanner_type": scanner_type,
			"active_rules": active_rules,
			"active_plugins": active_plugins,
		}
		bodyBytes, err := json.Marshal(bodyData)
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
// Getscannerrecommendations GET /scanner/recommendations
func (c *ScannerIntegrationClient) Getscannerrecommendations(scanner_type *string, current_rules *[]string, current_plugins *[]string) (interface{}, error) {
		url := fmt.Sprintf("/scanner/recommendations", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("scanner_type", fmt.Sprintf("%v", scanner_type))
		params.Add("current_rules", fmt.Sprintf("%v", current_rules))
		params.Add("current_plugins", fmt.Sprintf("%v", current_plugins))
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
