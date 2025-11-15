package tavo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// AiPerformanceQualityClient handles ai_performance_quality API calls
type AiPerformanceQualityClient struct {
	client *Client
}

// Getperformancemetrics GET /performance-metrics
func (c *AiPerformanceQualityClient) Getperformancemetrics(start_date *string, end_date *string, analysis_type *string) (interface{}, error) {
		url := fmt.Sprintf("/performance-metrics", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("start_date", fmt.Sprintf("%v", start_date))
		params.Add("end_date", fmt.Sprintf("%v", end_date))
		params.Add("analysis_type", fmt.Sprintf("%v", analysis_type))
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
// Getqualityreview{scan_id} GET /quality-review/{scan_id}
func (c *AiPerformanceQualityClient) Getqualityreview{scan_id}(scan_id string) (interface{}, error) {
		url := fmt.Sprintf("/quality-review/{scan_id}", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("scan_id", fmt.Sprintf("%v", scan_id))
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
