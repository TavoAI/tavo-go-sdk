package tavo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// AiResultsExportClient handles ai_results_export API calls
type AiResultsExportClient struct {
	client *Client
}

// Getresults GET /results
func (c *AiResultsExportClient) Getresults(skip *float64, limit *float64, scan_id *string, analysis_type *string, severity *string, start_date *string, end_date *string) (interface{}, error) {
		url := fmt.Sprintf("/results", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("skip", fmt.Sprintf("%v", skip))
		params.Add("limit", fmt.Sprintf("%v", limit))
		params.Add("scan_id", fmt.Sprintf("%v", scan_id))
		params.Add("analysis_type", fmt.Sprintf("%v", analysis_type))
		params.Add("severity", fmt.Sprintf("%v", severity))
		params.Add("start_date", fmt.Sprintf("%v", start_date))
		params.Add("end_date", fmt.Sprintf("%v", end_date))
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
// Getresultsexport GET /results/export
func (c *AiResultsExportClient) Getresultsexport(format *string, scan_id *string, analysis_type *string, start_date *string, end_date *string) (interface{}, error) {
		url := fmt.Sprintf("/results/export", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("format", fmt.Sprintf("%v", format))
		params.Add("scan_id", fmt.Sprintf("%v", scan_id))
		params.Add("analysis_type", fmt.Sprintf("%v", analysis_type))
		params.Add("start_date", fmt.Sprintf("%v", start_date))
		params.Add("end_date", fmt.Sprintf("%v", end_date))
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
