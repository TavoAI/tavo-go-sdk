package tavo

import (
	"context"
	"fmt"
	"net/url"
)

// ScannerOperations handles scanner integration operations for CLI tools and scanners
type ScannerOperations struct {
	client *Client
}

// DiscoverRules discovers rules optimized for scanner types
func (s *ScannerOperations) DiscoverRules(scannerType, language, category *string) (map[string]interface{}, error) {
	params := url.Values{}
	if scannerType != nil {
		params.Add("scanner_type", *scannerType)
	}
	if language != nil {
		params.Add("language", *language)
	}
	if category != nil {
		params.Add("category", *category)
	}

	path := "/scanner/rules/discover"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	return s.client.makeRequest("GET", path, nil)
}

// GetBundleRules gets rules from a specific bundle
func (s *ScannerOperations) GetBundleRules(bundleID string) (map[string]interface{}, error) {
	return s.client.makeRequest("GET", fmt.Sprintf("/scanner/rules/bundle/%s/rules", bundleID), nil)
}

// TrackBundleUsage tracks bundle usage by scanners
func (s *ScannerOperations) TrackBundleUsage(bundleID string, usageData map[string]interface{}) (map[string]interface{}, error) {
	if usageData == nil {
		usageData = make(map[string]interface{})
	}
	return s.client.makeRequest("POST", fmt.Sprintf("/scanner/rules/bundle/%s/use", bundleID), usageData)
}

// DiscoverPlugins discovers plugins optimized for scanner types
func (s *ScannerOperations) DiscoverPlugins(scannerType, language, category *string) (map[string]interface{}, error) {
	params := url.Values{}
	if scannerType != nil {
		params.Add("scanner_type", *scannerType)
	}
	if language != nil {
		params.Add("language", *language)
	}
	if category != nil {
		params.Add("category", *category)
	}

	path := "/scanner/plugins/discover"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	return s.client.makeRequest("GET", path, nil)
}

// GetPluginConfig gets plugin configuration for scanner use
func (s *ScannerOperations) GetPluginConfig(pluginID string) (map[string]interface{}, error) {
	return s.client.makeRequest("GET", fmt.Sprintf("/scanner/plugins/%s/config", pluginID), nil)
}

// GetRecommendations gets AI-powered rule/plugin recommendations
func (s *ScannerOperations) GetRecommendations(language, scannerType *string, currentRules, currentPlugins []string) (map[string]interface{}, error) {
	params := url.Values{}
	if language != nil {
		params.Add("language", *language)
	}
	if scannerType != nil {
		params.Add("scannerType", *scannerType)
	}
	for _, rule := range currentRules {
		params.Add("currentRules", rule)
	}
	for _, plugin := range currentPlugins {
		params.Add("currentPlugins", plugin)
	}

	path := "/scanner/recommendations"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	return s.client.makeRequest("GET", path, nil)
}

// SendHeartbeat sends scanner heartbeat for tracking
func (s *ScannerOperations) SendHeartbeat(heartbeatData map[string]interface{}) (map[string]interface{}, error) {
	return s.client.makeRequest("POST", "/scanner/heartbeat", heartbeatData)
}

// DiscoverRulesAsync discovers rules asynchronously
func (s *ScannerOperations) DiscoverRulesAsync(ctx context.Context, scannerType, language, category *string) (<-chan map[string]interface{}, <-chan error) {
	params := url.Values{}
	if scannerType != nil {
		params.Add("scanner_type", *scannerType)
	}
	if language != nil {
		params.Add("language", *language)
	}
	if category != nil {
		params.Add("category", *category)
	}

	path := "/scanner/rules/discover"
	if len(params) > 0 {
		path += "?" + params.Encode()
	}

	return s.client.makeRequestAsync(ctx, "GET", path, nil)
}

// SendHeartbeatAsync sends scanner heartbeat asynchronously
func (s *ScannerOperations) SendHeartbeatAsync(ctx context.Context, heartbeatData map[string]interface{}) (<-chan map[string]interface{}, <-chan error) {
	return s.client.makeRequestAsync(ctx, "POST", "/scanner/heartbeat", heartbeatData)
}
