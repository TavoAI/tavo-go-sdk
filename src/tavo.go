// Package tavo provides a client for the Tavo AI API
package tavo

import (
	"net/http"
)

// Client is the main client for interacting with Tavo AI API
type Client struct {
	httpClient *http.Client
	baseURL    string

	// Authentication
	apiKey      string
	deviceToken string

	deviceAuth *DeviceAuthClient
	scans *ScansClient
	scanManagement *ScanManagementClient
	scanTools *ScanToolsClient
	scanRules *ScanRulesClient
	scanSchedules *ScanSchedulesClient
	scanBulkOperations *ScanBulkOperationsClient
	scannerIntegration *ScannerIntegrationClient
	aiAnalysis *AiAnalysisClient
	aiAnalysisCore *AiAnalysisCoreClient
	aiBulkOperations *AiBulkOperationsClient
	aiPerformanceQuality *AiPerformanceQualityClient
	aiResultsExport *AiResultsExportClient
	aiRiskCompliance *AiRiskComplianceClient
	registry *RegistryClient
	pluginExecution *PluginExecutionClient
	pluginMarketplace *PluginMarketplaceClient
	rules *RulesClient
	codeSubmission *CodeSubmissionClient
	repositories *RepositoriesClient
	repositoryConnections *RepositoryConnectionsClient
	repositoryProviders *RepositoryProvidersClient
	repositoryWebhooks *RepositoryWebhooksClient
	jobs *JobsClient
	health *HealthClient
	websockets *WebsocketsClient
}

// NewClient creates a new Tavo API client
func NewClient(apiKey, deviceToken, baseUrl string) *Client {
	if baseUrl == "" {
		baseUrl = "https://api.tavo.ai"
	}

	httpClient := &http.Client{}

	client := &Client{
		httpClient:  httpClient,
		baseURL:     baseUrl,
		apiKey:      apiKey,
		deviceToken: deviceToken,
	}

	// Initialize endpoint clients
		client.deviceAuth = &DeviceAuthClient{client: client}
		client.scans = &ScansClient{client: client}
		client.scanManagement = &ScanManagementClient{client: client}
		client.scanTools = &ScanToolsClient{client: client}
		client.scanRules = &ScanRulesClient{client: client}
		client.scanSchedules = &ScanSchedulesClient{client: client}
		client.scanBulkOperations = &ScanBulkOperationsClient{client: client}
		client.scannerIntegration = &ScannerIntegrationClient{client: client}
		client.aiAnalysis = &AiAnalysisClient{client: client}
		client.aiAnalysisCore = &AiAnalysisCoreClient{client: client}
		client.aiBulkOperations = &AiBulkOperationsClient{client: client}
		client.aiPerformanceQuality = &AiPerformanceQualityClient{client: client}
		client.aiResultsExport = &AiResultsExportClient{client: client}
		client.aiRiskCompliance = &AiRiskComplianceClient{client: client}
		client.registry = &RegistryClient{client: client}
		client.pluginExecution = &PluginExecutionClient{client: client}
		client.pluginMarketplace = &PluginMarketplaceClient{client: client}
		client.rules = &RulesClient{client: client}
		client.codeSubmission = &CodeSubmissionClient{client: client}
		client.repositories = &RepositoriesClient{client: client}
		client.repositoryConnections = &RepositoryConnectionsClient{client: client}
		client.repositoryProviders = &RepositoryProvidersClient{client: client}
		client.repositoryWebhooks = &RepositoryWebhooksClient{client: client}
		client.jobs = &JobsClient{client: client}
		client.health = &HealthClient{client: client}
		client.websockets = &WebsocketsClient{client: client}

	return client
}

// SetAPIKey updates the API key for authentication
func (c *Client) SetAPIKey(apiKey string) {
	c.apiKey = apiKey
	c.deviceToken = ""
}

// SetDeviceToken updates the device token for authentication
func (c *Client) SetDeviceToken(deviceToken string) {
	c.deviceToken = deviceToken
	c.apiKey = ""
}
