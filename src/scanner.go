// Package scanner provides a wrapper for executing tavo-scanner as a subprocess
package scanner

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

// ScannerConfig represents configuration for tavo-scanner execution
type ScannerConfig struct {
	// Path to tavo-scanner binary
	ScannerPath string

	// List of plugins to use
	Plugins []string

	// Plugin-specific configuration
	PluginConfig map[string]interface{}

	// Path to custom rules file
	RulesPath string

	// Custom rules configuration
	CustomRules map[string]interface{}

	// Execution timeout in seconds
	Timeout int

	// Working directory for execution
	WorkingDirectory string

	// Output format
	OutputFormat string

	// Output file path
	OutputFile string
}

// NewScannerConfig creates a new scanner configuration with defaults
func NewScannerConfig() *ScannerConfig {
	return &ScannerConfig{
		ScannerPath:       findScannerBinary(),
		Plugins:           []string{},
		PluginConfig:      make(map[string]interface{}),
		CustomRules:       make(map[string]interface{}),
		Timeout:           300,
		WorkingDirectory:  ".",
		OutputFormat:      "json",
	}
}

// findScannerBinary finds the tavo-scanner binary in common locations
func findScannerBinary() string {
	// Try relative to this package
	_, filename, _, _ := runtime.Caller(0)
	packageDir := filepath.Dir(filename)
	scannerPath := filepath.Join(packageDir, "..", "..", "..", "..", "tavo-cli", "bin", "tavo-scanner")

	if _, err := os.Stat(scannerPath); err == nil {
		return scannerPath
	}

	// Check PATH
	if scannerPath, err := exec.LookPath("tavo-scanner"); err == nil {
		return scannerPath
	}

	return ""
}

// ScanOptions represents options for scanner execution
type ScanOptions struct {
	// Static analysis enabled
	StaticAnalysis bool

	// Static analysis plugins
	StaticPlugins []string

	// Custom rules path
	StaticRules string

	// Dynamic testing enabled
	DynamicTesting bool

	// Dynamic testing plugins
	DynamicPlugins []string

	// Output format
	OutputFormat string

	// Output file path
	OutputFile string

	// Execution timeout
	Timeout int

	// Files to exclude
	ExcludePatterns []string

	// Files to include
	IncludePatterns []string
}

// NewScanOptions creates default scan options
func NewScanOptions() *ScanOptions {
	return &ScanOptions{
		StaticAnalysis:  true,
		OutputFormat:    "json",
		Timeout:         300,
		ExcludePatterns: []string{},
		IncludePatterns: []string{},
	}
}

// ScanResult represents the result from scanner execution
type ScanResult struct {
	// Execution status
	Status string `json:"status"`

	// Scan results
	Results []interface{} `json:"results,omitempty"`

	// Raw output
	Output string `json:"output,omitempty"`

	// Error message
	Error string `json:"error,omitempty"`
}

// TavoScanner wraps tavo-scanner execution
type TavoScanner struct {
	config *ScannerConfig
}

// NewTavoScanner creates a new scanner wrapper
func NewTavoScanner(config *ScannerConfig) *TavoScanner {
	if config == nil {
		config = NewScannerConfig()
	}
	return &TavoScanner{config: config}
}

// ScanDirectory scans a directory with tavo-scanner
func (s *TavoScanner) ScanDirectory(targetPath string, scanOptions *ScanOptions) (*ScanResult, error) {
	if s.config.ScannerPath == "" {
		return nil, fmt.Errorf("tavo-scanner binary not found. Please install tavo-cli or set ScannerPath")
	}

	// Merge configurations
	mergedConfig := &ScannerConfig{
		ScannerPath:      s.config.ScannerPath,
		Plugins:          make([]string, len(s.config.Plugins)),
		PluginConfig:     make(map[string]interface{}),
		RulesPath:        s.config.RulesPath,
		CustomRules:      make(map[string]interface{}),
		Timeout:          s.config.Timeout,
		WorkingDirectory: s.config.WorkingDirectory,
		OutputFormat:     s.config.OutputFormat,
		OutputFile:       s.config.OutputFile,
	}
	copy(mergedConfig.Plugins, s.config.Plugins)
	for k, v := range s.config.PluginConfig {
		mergedConfig.PluginConfig[k] = v
	}
	for k, v := range s.config.CustomRules {
		mergedConfig.CustomRules[k] = v
	}

	if scanOptions != nil {
		mergedConfig.Plugins = scanOptions.StaticPlugins
		mergedConfig.RulesPath = scanOptions.StaticRules
		mergedConfig.Timeout = scanOptions.Timeout
		mergedConfig.OutputFormat = scanOptions.OutputFormat
		mergedConfig.OutputFile = scanOptions.OutputFile
	}

	// Prepare command arguments
	args := []string{targetPath}

	// Add plugins
	for _, plugin := range mergedConfig.Plugins {
		args = append(args, "--plugin", plugin)
	}

	// Add rules
	if mergedConfig.RulesPath != "" {
		args = append(args, "--rules", mergedConfig.RulesPath)
	}

	// Add output options
	if mergedConfig.OutputFormat != "" {
		args = append(args, "--format", mergedConfig.OutputFormat)
	}

	if mergedConfig.OutputFile != "" {
		args = append(args, "--output", mergedConfig.OutputFile)
	}

	// Add timeout
	if mergedConfig.Timeout > 0 {
		args = append(args, "--timeout", fmt.Sprintf("%d", mergedConfig.Timeout))
	}

	return s.executeScanner(args, mergedConfig.WorkingDirectory)
}

// ScanWithPlugins scans with specific plugins
func (s *TavoScanner) ScanWithPlugins(targetPath string, plugins []string) (*ScanResult, error) {
	options := NewScanOptions()
	options.StaticPlugins = plugins
	return s.ScanDirectory(targetPath, options)
}

// ScanWithRules scans with custom rules
func (s *TavoScanner) ScanWithRules(targetPath, rulesPath string) (*ScanResult, error) {
	options := NewScanOptions()
	options.StaticRules = rulesPath
	return s.ScanDirectory(targetPath, options)
}

// executeScanner executes the scanner subprocess
func (s *TavoScanner) executeScanner(args []string, workingDirectory string) (*ScanResult, error) {
	cmd := exec.Command(s.config.ScannerPath, args...)
	cmd.Dir = workingDirectory

	// Set timeout
	done := make(chan bool)
	var result *ScanResult
	var execErr error

	go func() {
		output, err := cmd.CombinedOutput()
		if err != nil {
			execErr = err
			result = &ScanResult{
				Status: "error",
				Error:  string(output),
			}
			if result.Error == "" {
				result.Error = fmt.Sprintf("Scanner exited with code %d", cmd.ProcessState.ExitCode())
			}
		} else {
			outputStr := string(output)
			result = &ScanResult{Status: "success"}

			if outputStr == "" {
				result.Results = []interface{}{}
			} else {
				// Try to parse as JSON
				var results []interface{}
				if err := json.Unmarshal([]byte(outputStr), &results); err == nil {
					result.Results = results
				} else {
					result.Output = outputStr
				}
			}
		}
		done <- true
	}()

	// Wait for completion or timeout
	select {
	case <-done:
		return result, execErr
	case <-time.After(time.Duration(s.config.Timeout) * time.Second):
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
		return nil, fmt.Errorf("scanner timed out after %d seconds", s.config.Timeout)
	}
}

// CreatePluginConfig creates a temporary plugin configuration file
func (s *TavoScanner) CreatePluginConfig(pluginName string, config map[string]interface{}) (string, error) {
	tempFile, err := os.CreateTemp("", fmt.Sprintf("tavo-plugin-%s-*.json", pluginName))
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	encoder := json.NewEncoder(tempFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		return "", err
	}

	return tempFile.Name(), nil
}

// CreateRulesFile creates a temporary rules file
func (s *TavoScanner) CreateRulesFile(rules map[string]interface{}) (string, error) {
	tempFile, err := os.CreateTemp("", "tavo-rules-*.json")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	encoder := json.NewEncoder(tempFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(rules); err != nil {
		return "", err
	}

	return tempFile.Name(), nil
}
