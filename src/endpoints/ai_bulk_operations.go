package tavo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// AiBulkOperationsClient handles ai_bulk_operations API calls
type AiBulkOperationsClient struct {
	client *Client
}

// Deletebulkdelete DELETE /bulk/delete
func (c *AiBulkOperationsClient) Deletebulkdelete(analysis_ids *[]string) (interface{}, error) {
		url := fmt.Sprintf("/bulk/delete", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyBytes, err := json.Marshal(analysis_ids)
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
// Putbulkupdatestatus PUT /bulk/update-status
func (c *AiBulkOperationsClient) Putbulkupdatestatus(analysis_updates *[]interface{}) (interface{}, error) {
		url := fmt.Sprintf("/bulk/update-status", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyBytes, err := json.Marshal(analysis_updates)
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
// Getbulkexport GET /bulk/export
func (c *AiBulkOperationsClient) Getbulkexport(analysis_ids *[]string, export_format *string) (interface{}, error) {
		url := fmt.Sprintf("/bulk/export", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("analysis_ids", fmt.Sprintf("%v", analysis_ids))
		params.Add("export_format", fmt.Sprintf("%v", export_format))
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
