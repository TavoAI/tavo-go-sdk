package tavo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// AiAnalysisClient handles ai_analysis API calls
type AiAnalysisClient struct {
	client *Client
}

// Postanalyze{scan_id} POST /analyze/{scan_id}
func (c *AiAnalysisClient) Postanalyze{scan_id}(scan_id string, background_tasks interface{}) (interface{}, error) {
		url := fmt.Sprintf("/analyze/{scan_id}", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyData := map[string]interface{}{
			"scan_id": scan_id,
			"background_tasks": background_tasks,
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
// Postclassify{scan_id} POST /classify/{scan_id}
func (c *AiAnalysisClient) Postclassify{scan_id}(scan_id string, background_tasks interface{}) (interface{}, error) {
		url := fmt.Sprintf("/classify/{scan_id}", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyData := map[string]interface{}{
			"scan_id": scan_id,
			"background_tasks": background_tasks,
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
// Postriskscore{scan_id} POST /risk-score/{scan_id}
func (c *AiAnalysisClient) Postriskscore{scan_id}(scan_id string) (interface{}, error) {
		url := fmt.Sprintf("/risk-score/{scan_id}", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyBytes, err := json.Marshal(scan_id)
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
// Postcompliance{scan_id} POST /compliance/{scan_id}
func (c *AiAnalysisClient) Postcompliance{scan_id}(scan_id float64, framework *string) (interface{}, error) {
		url := fmt.Sprintf("/compliance/{scan_id}", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyData := map[string]interface{}{
			"scan_id": scan_id,
			"framework": framework,
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
// Postpredictive{scan_id} POST /predictive/{scan_id}
func (c *AiAnalysisClient) Postpredictive{scan_id}(scan_id string) (interface{}, error) {
		url := fmt.Sprintf("/predictive/{scan_id}", )
		fullURL := c.client.baseURL + "/api/v1" + url
		bodyBytes, err := json.Marshal(scan_id)
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
// Getfixsuggestions GET /fix-suggestions
func (c *AiAnalysisClient) Getfixsuggestions(search *string, status *string, severity *string, analysis_type *string, limit *float64, offset *float64) (interface{}, error) {
		url := fmt.Sprintf("/fix-suggestions", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("search", fmt.Sprintf("%v", search))
		params.Add("status", fmt.Sprintf("%v", status))
		params.Add("severity", fmt.Sprintf("%v", severity))
		params.Add("analysis_type", fmt.Sprintf("%v", analysis_type))
		params.Add("limit", fmt.Sprintf("%v", limit))
		params.Add("offset", fmt.Sprintf("%v", offset))
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
// Getpredictive GET /predictive
func (c *AiAnalysisClient) Getpredictive(time_horizon *string, severity *string, prediction_type *string, analysis_type *string) (interface{}, error) {
		url := fmt.Sprintf("/predictive", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("time_horizon", fmt.Sprintf("%v", time_horizon))
		params.Add("severity", fmt.Sprintf("%v", severity))
		params.Add("prediction_type", fmt.Sprintf("%v", prediction_type))
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
// Getcompliance GET /compliance
func (c *AiAnalysisClient) Getcompliance(framework *string, status *string, risk_level *string, category *string) (interface{}, error) {
		url := fmt.Sprintf("/compliance", )
		fullURL := c.client.baseURL + "/api/v1" + url
		params := url.Values{}
		params.Add("framework", fmt.Sprintf("%v", framework))
		params.Add("status", fmt.Sprintf("%v", status))
		params.Add("risk_level", fmt.Sprintf("%v", risk_level))
		params.Add("category", fmt.Sprintf("%v", category))
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
