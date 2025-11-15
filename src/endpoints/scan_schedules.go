package tavo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// ScanSchedulesClient handles scan_schedules API calls
type ScanSchedulesClient struct {
	client *Client
}

// PostRoot POST /
func (c *ScanSchedulesClient) PostRoot(schedule_in interface{}) (interface{}, error) {
		url := fmt.Sprintf("/", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyBytes, err := json.Marshal(schedule_in)
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
// Getrepository{repository_id} GET /repository/{repository_id}
func (c *ScanSchedulesClient) Getrepository{repository_id}(repository_id string) (interface{}, error) {
		url := fmt.Sprintf("/repository/{repository_id}", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("repository_id", fmt.Sprintf("%v", repository_id))
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
// Get{schedule_id} GET /{schedule_id}
func (c *ScanSchedulesClient) Get{schedule_id}(schedule_id string) (interface{}, error) {
		url := fmt.Sprintf("/{schedule_id}", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("schedule_id", fmt.Sprintf("%v", schedule_id))
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
// Put{schedule_id} PUT /{schedule_id}
func (c *ScanSchedulesClient) Put{schedule_id}(schedule_id string, schedule_update interface{}) (interface{}, error) {
		url := fmt.Sprintf("/{schedule_id}", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyData := map[string]interface{}{
			"schedule_id": schedule_id,
			"schedule_update": schedule_update,
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
// Delete{schedule_id} DELETE /{schedule_id}
func (c *ScanSchedulesClient) Delete{schedule_id}(schedule_id string) (interface{}, error) {
		url := fmt.Sprintf("/{schedule_id}", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyBytes, err := json.Marshal(schedule_id)
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
