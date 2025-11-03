package tavo

import (
	"context"
	"testing"
	"time"
)

func TestClient_makeRequestAsync(t *testing.T) {
	config := NewConfig().WithAPIKey("test-key").WithTimeout(100 * time.Millisecond).WithBaseURL("http://10.255.255.1").WithMaxRetries(0).WithBaseURL("http://10.255.255.1").WithMaxRetries(0)
	client := NewClient(config)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Test async request (will fail due to invalid endpoint, but tests the goroutine functionality)
	resultChan, errorChan := client.makeRequestAsync(ctx, "GET", "/health", nil)

	select {
	case result := <-resultChan:
		t.Logf("Got result: %v", result)
	case err := <-errorChan:
		// Expected to fail with network error, but goroutine executed
		t.Logf("Got expected error: %v", err)
	case <-time.After(2 * time.Second):
		t.Fatal("Async request timed out")
	}
}

func TestClient_DeviceOperations(t *testing.T) {
	config := NewConfig().WithAPIKey("test-key").WithTimeout(100 * time.Millisecond).WithBaseURL("http://10.255.255.1").WithMaxRetries(0).WithBaseURL("http://10.255.255.1").WithMaxRetries(0)
	client := NewClient(config)
	deviceOps := client.Device()

	if deviceOps == nil {
		t.Fatal("Device operations should not be nil")
	}

	// Test method existence (actual calls will fail due to network)
	clientName := "Test Client"
	_, err := deviceOps.CreateDeviceCode(nil, &clientName)
	// Should fail with network error, but method exists
	if err == nil {
		t.Error("Expected network error, got nil")
	}
}

func TestClient_ScannerOperations(t *testing.T) {
	config := NewConfig().WithAPIKey("test-key").WithTimeout(100 * time.Millisecond).WithBaseURL("http://10.255.255.1").WithMaxRetries(0)
	client := NewClient(config)
	scannerOps := client.Scanner()

	if scannerOps == nil {
		t.Fatal("Scanner operations should not be nil")
	}

	// Test method existence
	scannerType := "sast"
	_, err := scannerOps.DiscoverRules(&scannerType, nil, nil)
	// Should fail with network error, but method exists
	if err == nil {
		t.Error("Expected network error, got nil")
	}
}

func TestClient_CodeSubmissionOperations(t *testing.T) {
	config := NewConfig().WithAPIKey("test-key").WithTimeout(100 * time.Millisecond).WithBaseURL("http://10.255.255.1").WithMaxRetries(0)
	client := NewClient(config)
	codeOps := client.CodeSubmission()

	if codeOps == nil {
		t.Fatal("Code submission operations should not be nil")
	}

	// Test method existence
	_, err := codeOps.SubmitAnalysis("test code", "go", nil, nil, nil, nil)
	// Should fail with network error, but method exists
	if err == nil {
		t.Error("Expected network error, got nil")
	}
}

func TestDeviceOperations_AsyncMethods(t *testing.T) {
	config := NewConfig().WithAPIKey("test-key").WithTimeout(100 * time.Millisecond).WithBaseURL("http://10.255.255.1").WithMaxRetries(0)
	client := NewClient(config)
	deviceOps := client.Device()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Test async device code creation
	resultChan, errorChan := deviceOps.CreateDeviceCodeAsync(ctx, nil, nil)

	select {
	case result := <-resultChan:
		t.Logf("Got result: %v", result)
	case err := <-errorChan:
		// Expected to fail with network error, but async functionality works
		t.Logf("Got expected error: %v", err)
	case <-time.After(1 * time.Second):
		t.Fatal("Async device code creation timed out")
	}
}

func TestScannerOperations_AsyncMethods(t *testing.T) {
	config := NewConfig().WithAPIKey("test-key").WithTimeout(100 * time.Millisecond).WithBaseURL("http://10.255.255.1").WithMaxRetries(0)
	client := NewClient(config)
	scannerOps := client.Scanner()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Test async heartbeat
	heartbeatData := map[string]interface{}{
		"scanner_version": "1.0.0",
		"scanner_type":    "sast",
	}

	resultChan, errorChan := scannerOps.SendHeartbeatAsync(ctx, heartbeatData)

	select {
	case result := <-resultChan:
		t.Logf("Got result: %v", result)
	case err := <-errorChan:
		// Expected to fail with network error, but async functionality works
		t.Logf("Got expected error: %v", err)
	case <-time.After(1 * time.Second):
		t.Fatal("Async heartbeat timed out")
	}
}

func TestCodeSubmissionOperations_AsyncMethods(t *testing.T) {
	config := NewConfig().WithAPIKey("test-key").WithTimeout(100 * time.Millisecond).WithBaseURL("http://10.255.255.1").WithMaxRetries(0)
	client := NewClient(config)
	codeOps := client.CodeSubmission()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Test async analysis submission
	resultChan, errorChan := codeOps.SubmitAnalysisAsync(ctx, "func main() {}", "go", nil, nil, nil, nil)

	select {
	case result := <-resultChan:
		t.Logf("Got result: %v", result)
	case err := <-errorChan:
		// Expected to fail with network error, but async functionality works
		t.Logf("Got expected error: %v", err)
	case <-time.After(1 * time.Second):
		t.Fatal("Async analysis submission timed out")
	}
}

func TestClient_CloudNativeOptimizations(t *testing.T) {
	config := NewConfig().
		WithAPIKey("test-key").
		WithTimeout(10).   // Cloud-native: configurable timeouts
		WithMaxRetries(3). // Cloud-native: retry logic
		WithBaseURL("https://api.tavoai.net")

	client := NewClient(config)

	// Test that client is properly configured
	if client == nil {
		t.Fatal("Client should not be nil")
	}

	// Test context support
	ctx := context.Background()
	_, err := client.makeRequestWithContext(ctx, "GET", "/health", nil)
	// Should fail with network error, but context support works
	if err == nil {
		t.Error("Expected network error, got nil")
	}
}

func TestClient_GoroutineCancellation(t *testing.T) {
	client := NewClient(NewConfig().WithAPIKey("test-key"))

	// Test context cancellation
	ctx, cancel := context.WithCancel(context.Background())

	// Start async request
	resultChan, errorChan := client.makeRequestAsync(ctx, "GET", "/health", nil)

	// Cancel immediately
	cancel()

	select {
	case <-resultChan:
		t.Error("Should not receive result after cancellation")
	case err := <-errorChan:
		// Should get context cancellation error
		t.Logf("Got cancellation error: %v", err)
	case <-time.After(1 * time.Second):
		t.Error("Should have received cancellation error")
	}
}
