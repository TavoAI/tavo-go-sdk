package tavo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// AiRiskComplianceClient handles ai_risk_compliance API calls
type AiRiskComplianceClient struct {
	client *Client
}

// Getriskscores GET /risk-scores
func (c *AiRiskComplianceClient) Getriskscores(skip *float64, limit *float64, scan_id *string, min_score *float64, max_score *float64) (interface{}, error) {
		url := fmt.Sprintf("/risk-scores", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("skip", fmt.Sprintf("%v", skip))
		params.Add("limit", fmt.Sprintf("%v", limit))
		params.Add("scan_id", fmt.Sprintf("%v", scan_id))
		params.Add("min_score", fmt.Sprintf("%v", min_score))
		params.Add("max_score", fmt.Sprintf("%v", max_score))
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
// Getcompliancereports GET /compliance-reports
func (c *AiRiskComplianceClient) Getcompliancereports(skip *float64, limit *float64, scan_id *string, framework *string, status *string) (interface{}, error) {
		url := fmt.Sprintf("/compliance-reports", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("skip", fmt.Sprintf("%v", skip))
		params.Add("limit", fmt.Sprintf("%v", limit))
		params.Add("scan_id", fmt.Sprintf("%v", scan_id))
		params.Add("framework", fmt.Sprintf("%v", framework))
		params.Add("status", fmt.Sprintf("%v", status))
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
// Getpredictiveanalyses GET /predictive-analyses
func (c *AiRiskComplianceClient) Getpredictiveanalyses(skip *float64, limit *float64, scan_id *string, prediction_type *string, confidence_threshold *float64) (interface{}, error) {
		url := fmt.Sprintf("/predictive-analyses", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("skip", fmt.Sprintf("%v", skip))
		params.Add("limit", fmt.Sprintf("%v", limit))
		params.Add("scan_id", fmt.Sprintf("%v", scan_id))
		params.Add("prediction_type", fmt.Sprintf("%v", prediction_type))
		params.Add("confidence_threshold", fmt.Sprintf("%v", confidence_threshold))
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
