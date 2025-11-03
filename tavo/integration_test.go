package tavo

import (
	"context"
	"testing"
	"time"
)

// These tests require a running API server and proper credentials
// They are marked as integration tests and may be skipped in CI

func TestDeviceOperations_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// This would require a valid API key and running server
	t.Skip("Integration test - requires running API server")

	client := NewClient(NewConfig().WithAPIKey("valid-api-key"))
	deviceOps := client.Device()

	// Test creating device code
	clientName := "Test Integration"
	result, err := deviceOps.CreateDeviceCode(nil, &clientName)

	if err != nil {
		t.Fatalf("Failed to create device code: %v", err)
	}

	if result == nil {
		t.Fatal("Result should not be nil")
	}

	// Verify expected fields in response
	if _, exists := result["device_code"]; !exists {
		t.Error("Response should contain device_code")
	}
	if _, exists := result["user_code"]; !exists {
		t.Error("Response should contain user_code")
	}
	if _, exists := result["verification_uri"]; !exists {
		t.Error("Response should contain verification_uri")
	}
}

func TestScannerOperations_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	t.Skip("Integration test - requires running API server")

	client := NewClient(NewConfig().WithAPIKey("valid-api-key"))
	scannerOps := client.Scanner()

	// Test discovering rules
	scannerType := "sast"
	result, err := scannerOps.DiscoverRules(&scannerType, nil, nil)

	if err != nil {
		t.Fatalf("Failed to discover rules: %v", err)
	}

	if result == nil {
		t.Fatal("Result should not be nil")
	}
}

func TestCodeSubmissionOperations_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	t.Skip("Integration test - requires running API server")

	client := NewClient(NewConfig().WithAPIKey("valid-api-key"))
	codeOps := client.CodeSubmission()

	// Test submitting code analysis
	codeContent := `package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}`

	result, err := codeOps.SubmitAnalysis(codeContent, "go", nil, nil, nil, nil)

	if err != nil {
		t.Fatalf("Failed to submit analysis: %v", err)
	}

	if result == nil {
		t.Fatal("Result should not be nil")
	}

	// Verify expected fields
	if _, exists := result["scan_id"]; !exists {
		t.Error("Response should contain scan_id")
	}
	if status, exists := result["status"]; !exists || status != "accepted" {
		t.Error("Response should contain status=accepted")
	}
}

func TestConcurrentOperations_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	t.Skip("Integration test - requires running API server")

	client := NewClient(NewConfig().WithAPIKey("valid-api-key"))
	ctx := context.Background()

	// Test concurrent operations using goroutines
	results := make(chan string, 3)
	errors := make(chan error, 3)

	// Start multiple async operations
	go func() {
		deviceOps := client.Device()
		resultChan, errorChan := deviceOps.CreateDeviceCodeAsync(ctx, nil, nil)

		select {
		case result := <-resultChan:
			if scanID, ok := result["device_code"].(string); ok {
				results <- "device_code:" + scanID
			} else {
				results <- "device_code:success"
			}
		case err := <-errorChan:
			errors <- err
		}
	}()

	go func() {
		scannerOps := client.Scanner()
		resultChan, errorChan := scannerOps.DiscoverRulesAsync(ctx, nil, nil, nil)

		select {
		case <-resultChan:
			results <- "rules:success"
		case err := <-errorChan:
			errors <- err
		}
	}()

	go func() {
		codeOps := client.CodeSubmission()
		resultChan, errorChan := codeOps.SubmitAnalysisAsync(ctx, "test", "go", nil, nil, nil, nil)

		select {
		case result := <-resultChan:
			if scanID, ok := result["scan_id"].(string); ok {
				results <- "analysis:" + scanID
			} else {
				results <- "analysis:success"
			}
		case err := <-errorChan:
			errors <- err
		}
	}()

	// Wait for results with timeout
	timeout := time.After(30 * time.Second)
	received := 0

	for received < 3 {
		select {
		case result := <-results:
			t.Logf("Received result: %s", result)
			received++
		case err := <-errors:
			t.Errorf("Received error: %v", err)
			received++
		case <-timeout:
			t.Fatal("Test timed out waiting for concurrent operations")
		}
	}
}

func BenchmarkClient_makeRequest(b *testing.B) {
	client := NewClient(NewConfig().WithAPIKey("test-key"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// This will fail with network error, but tests the request setup overhead
		_, _ = client.makeRequest("GET", "/health", nil)
	}
}

func BenchmarkClient_makeRequestAsync(b *testing.B) {
	client := NewClient(NewConfig().WithAPIKey("test-key"))
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resultChan, _ := client.makeRequestAsync(ctx, "GET", "/health", nil)
		// Don't wait for result, just test goroutine creation overhead
		go func() {
			<-resultChan
		}()
	}
}

func BenchmarkDeviceOperations_CreateDeviceCode(b *testing.B) {
	client := NewClient(NewConfig().WithAPIKey("test-key"))
	deviceOps := client.Device()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// This will fail with network error, but tests the method call overhead
		_, _ = deviceOps.CreateDeviceCode(nil, nil)
	}
}
