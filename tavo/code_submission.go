package tavo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
)

// CodeSubmissionOperations handles code submission operations for CLI tools and scanners
type CodeSubmissionOperations struct {
	client *Client
}

// SubmitCode submits code files directly for scanning
func (c *CodeSubmissionOperations) SubmitCode(files [][]byte, filenames []string, repositoryName, branch, commitSha *string, scanConfig map[string]interface{}) (map[string]interface{}, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// Add files
	for i, fileContent := range files {
		filename := "file"
		if i < len(filenames) {
			filename = filenames[i]
		}

		fw, err := w.CreateFormFile("files", filename)
		if err != nil {
			return nil, fmt.Errorf("failed to create form file: %w", err)
		}

		if _, err := fw.Write(fileContent); err != nil {
			return nil, fmt.Errorf("failed to write file content: %w", err)
		}
	}

	// Add optional parameters
	if repositoryName != nil {
		if err := w.WriteField("repository_name", *repositoryName); err != nil {
			return nil, fmt.Errorf("failed to write repository_name: %w", err)
		}
	}
	if branch != nil {
		if err := w.WriteField("branch", *branch); err != nil {
			return nil, fmt.Errorf("failed to write branch: %w", err)
		}
	}
	if commitSha != nil {
		if err := w.WriteField("commit_sha", *commitSha); err != nil {
			return nil, fmt.Errorf("failed to write commit_sha: %w", err)
		}
	}
	if scanConfig != nil {
		// This would need JSON encoding, simplified for now
		if err := w.WriteField("scan_config", fmt.Sprintf("%v", scanConfig)); err != nil {
			return nil, fmt.Errorf("failed to write scan_config: %w", err)
		}
	}

	w.Close()

	// Create a custom request for multipart data
	req := c.client.httpClient.R().
		SetHeader("Content-Type", w.FormDataContentType()).
		SetBody(b.Bytes())

	resp, err := req.Post("/code/submit/code")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, &TavoError{
			StatusCode: resp.StatusCode(),
			Message:    resp.Status(),
		}
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return result, nil
}

// SubmitRepository submits repository snapshot for scanning
func (c *CodeSubmissionOperations) SubmitRepository(repositoryURL string, snapshotData map[string]interface{}, branch, commitSha *string, scanConfig map[string]interface{}) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"repository_url": repositoryURL,
		"snapshot_data":  snapshotData,
	}
	if branch != nil {
		data["branch"] = *branch
	}
	if commitSha != nil {
		data["commit_sha"] = *commitSha
	}
	if scanConfig != nil {
		data["scan_config"] = scanConfig
	}

	return c.client.makeRequest("POST", "/code/submit/repository", data)
}

// SubmitAnalysis submits code snippet for targeted analysis
func (c *CodeSubmissionOperations) SubmitAnalysis(codeContent, language string, analysisType *string, rules, plugins []string, contextData map[string]interface{}) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"code_content": codeContent,
		"language":     language,
	}
	if analysisType != nil {
		data["analysis_type"] = *analysisType
	}
	if len(rules) > 0 {
		data["rules"] = rules
	}
	if len(plugins) > 0 {
		data["plugins"] = plugins
	}
	if contextData != nil {
		data["context"] = contextData
	}

	return c.client.makeRequest("POST", "/code/submit/analysis", data)
}

// GetScanStatus gets scan status (CLI-optimized)
func (c *CodeSubmissionOperations) GetScanStatus(scanID string) (map[string]interface{}, error) {
	return c.client.makeRequest("GET", fmt.Sprintf("/code/scans/%s/status", scanID), nil)
}

// GetScanResults gets scan results summary (CLI-optimized)
func (c *CodeSubmissionOperations) GetScanResults(scanID string) (map[string]interface{}, error) {
	return c.client.makeRequest("GET", fmt.Sprintf("/code/scans/%s/results/summary", scanID), nil)
}

// SubmitCodeAsync submits code files asynchronously
func (c *CodeSubmissionOperations) SubmitCodeAsync(ctx context.Context, files [][]byte, filenames []string, repositoryName, branch, commitSha *string, scanConfig map[string]interface{}) (<-chan map[string]interface{}, <-chan error) {
	resultChan := make(chan map[string]interface{}, 1)
	errorChan := make(chan error, 1)

	go func() {
		defer close(resultChan)
		defer close(errorChan)

		result, err := c.SubmitCode(files, filenames, repositoryName, branch, commitSha, scanConfig)
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- result
	}()

	return resultChan, errorChan
}

// SubmitAnalysisAsync submits code analysis asynchronously
func (c *CodeSubmissionOperations) SubmitAnalysisAsync(ctx context.Context, codeContent, language string, analysisType *string, rules, plugins []string, contextData map[string]interface{}) (<-chan map[string]interface{}, <-chan error) {
	data := map[string]interface{}{
		"code_content": codeContent,
		"language":     language,
	}
	if analysisType != nil {
		data["analysis_type"] = *analysisType
	}
	if len(rules) > 0 {
		data["rules"] = rules
	}
	if len(plugins) > 0 {
		data["plugins"] = plugins
	}
	if contextData != nil {
		data["context"] = contextData
	}

	return c.client.makeRequestAsync(ctx, "POST", "/code/submit/analysis", data)
}
