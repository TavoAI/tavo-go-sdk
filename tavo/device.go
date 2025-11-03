package tavo

import (
	"context"
	"fmt"
)

// DeviceOperations handles device authentication operations for CLI tools
type DeviceOperations struct {
	client *Client
}

// CreateDeviceCode creates a device code for authentication
func (d *DeviceOperations) CreateDeviceCode(clientID, clientName *string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	if clientID != nil {
		data["client_id"] = *clientID
	}
	if clientName != nil {
		data["client_name"] = *clientName
	}
	return d.client.makeRequest("POST", "/device/code", data)
}

// CreateDeviceCodeForCli creates a CLI-optimized device code for authentication
func (d *DeviceOperations) CreateDeviceCodeForCli(clientID, clientName *string) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	if clientID != nil {
		data["client_id"] = *clientID
	}
	data["client_name"] = "Tavo CLI"
	if clientName != nil {
		data["client_name"] = *clientName
	}
	return d.client.makeRequest("POST", "/device/code/cli", data)
}

// PollDeviceToken polls for a device token
func (d *DeviceOperations) PollDeviceToken(deviceCode string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"device_code": deviceCode,
	}
	return d.client.makeRequest("POST", "/device/token", data)
}

// GetDeviceCodeStatus gets device code status (lightweight polling for CLI)
func (d *DeviceOperations) GetDeviceCodeStatus(deviceCode string) (map[string]interface{}, error) {
	return d.client.makeRequest("GET", fmt.Sprintf("/device/code/%s/status", deviceCode), nil)
}

// GetUsageWarnings gets usage warnings and limits for CLI tools
func (d *DeviceOperations) GetUsageWarnings() (map[string]interface{}, error) {
	return d.client.makeRequest("GET", "/device/usage/warnings", nil)
}

// GetLimits gets current limits and quotas for CLI tools
func (d *DeviceOperations) GetLimits() (map[string]interface{}, error) {
	return d.client.makeRequest("GET", "/device/limits", nil)
}

// CreateDeviceCodeAsync creates a device code asynchronously
func (d *DeviceOperations) CreateDeviceCodeAsync(ctx context.Context, clientID, clientName *string) (<-chan map[string]interface{}, <-chan error) {
	data := make(map[string]interface{})
	if clientID != nil {
		data["client_id"] = *clientID
	}
	if clientName != nil {
		data["client_name"] = *clientName
	}
	return d.client.makeRequestAsync(ctx, "POST", "/device/code", data)
}

// PollDeviceTokenAsync polls for a device token asynchronously
func (d *DeviceOperations) PollDeviceTokenAsync(ctx context.Context, deviceCode string) (<-chan map[string]interface{}, <-chan error) {
	data := map[string]interface{}{
		"device_code": deviceCode,
	}
	return d.client.makeRequestAsync(ctx, "POST", "/device/token", data)
}
