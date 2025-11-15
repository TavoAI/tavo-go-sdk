// Package sdk provides the main entry point for Tavo AI SDK
package sdk

import (
	"github.com/tavo-ai/sdk-go/endpoints"
	"github.com/tavo-ai/sdk-go/scanner"
)

// Version represents the SDK version
const Version = "0.1.0"

// CreateClient creates a new API client instance
func CreateClient() *endpoints.TavoClient {
	return endpoints.NewTavoClient(nil)
}

// CreateClientWithAuth creates a new API client instance with authentication
func CreateClientWithAuth(apiKey, jwtToken, sessionToken string) *endpoints.TavoClient {
	config := &endpoints.TavoConfig{
		APIKey:       apiKey,
		JWTToken:     jwtToken,
		SessionToken: sessionToken,
	}
	return endpoints.NewTavoClient(config)
}

// CreateScanner creates a new scanner instance
func CreateScanner() *scanner.TavoScanner {
	return scanner.NewTavoScanner(nil)
}

// CreateScannerWithConfig creates a new scanner instance with configuration
func CreateScannerWithConfig(config *scanner.ScannerConfig) *scanner.TavoScanner {
	return scanner.NewTavoScanner(config)
}

// CreateScannerWithPlugins creates a scanner with specific plugins
func CreateScannerWithPlugins(plugins ...string) *scanner.TavoScanner {
	config := scanner.NewScannerConfig()
	config.Plugins = plugins
	return scanner.NewTavoScanner(config)
}

// CreateScannerWithRules creates a scanner with custom rules
func CreateScannerWithRules(rulesPath string) *scanner.TavoScanner {
	config := scanner.NewScannerConfig()
	config.RulesPath = rulesPath
	return scanner.NewTavoScanner(config)
}
