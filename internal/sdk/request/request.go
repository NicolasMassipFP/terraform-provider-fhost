// Package request provides HTTP request utilities for the SMC provider.
package request

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-smc/internal/sdk/auth"
)

// Options defines options for the generic request
type Options struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    []byte
	Timeout time.Duration
	Context context.Context
	Auth    *auth.Auth
}

// ResponseData contains only the body, headers, and status code
// for easier consumption in callers
type ResponseData struct {
	Body       []byte
	Headers    http.Header
	StatusCode int
	ETag       string
	Location   string
}

// DoRequest performs a generic HTTP request and returns ResponseData and error
func DoRequest(opts Options) (*ResponseData, error) {
	ctx := opts.Context
	if ctx == nil {
		ctx = context.Background()
	}

	var body io.Reader
	if opts.Body != nil {
		body = bytes.NewReader(opts.Body)
	}

	req, err := http.NewRequestWithContext(ctx, opts.Method, opts.URL, body)
	if err != nil {
		return nil, err
	}

	for k, v := range opts.Headers {
		req.Header.Set(k, v)
	}

	// Create HTTP client with optional TLS configuration
	client := &http.Client{}
	if opts.Timeout > 0 {
		client.Timeout = opts.Timeout
	}

	// Configure HTTPS if TrustedCert is provided
	if opts.Auth != nil && opts.Auth.TrustedCert != "" {
		tlsConfig, err := createTLSConfig(opts.Auth.TrustedCert, opts.Auth.Insecure)
		if err != nil {
			return nil, fmt.Errorf("failed to create TLS config: %w", err)
		}

		transport := &http.Transport{
			TLSClientConfig: tlsConfig,
		}
		client.Transport = transport
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	etag := resp.Header.Get("ETag")
	if etag != "" {
		etag = strings.ReplaceAll(etag, "\"", "")
	} else {
		etag = ""
	}

	location := resp.Header.Get("Location")

	fmt.Printf("Response StatusCode: %d\n", resp.StatusCode)
	fmt.Printf("Response Headers: %v\n", resp.Header)
	fmt.Printf("Response Body: %s\n", string(respBody))
	fmt.Printf("Response ETag: %s\n", etag)
	fmt.Printf("Response Location: %s\n", location)

	return &ResponseData{
		Body:       respBody,
		Headers:    resp.Header,
		StatusCode: resp.StatusCode,
		ETag:       etag,
		Location:   location,
	}, nil
}

// createTLSConfig creates a TLS configuration with trusted certificate
func createTLSConfig(trustedCert string, insecure *bool) (*tls.Config, error) {
	tlsConfig := &tls.Config{}

	// If insecure is explicitly set to true, skip certificate verification
	if insecure != nil && *insecure {
		tlsConfig.InsecureSkipVerify = true
		return tlsConfig, nil
	}

	// Create certificate pool and add the trusted certificate content
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM([]byte(trustedCert)) {
		return nil, fmt.Errorf("failed to append trusted certificate")
	}

	tlsConfig.RootCAs = caCertPool
	tlsConfig.InsecureSkipVerify = false

	return tlsConfig, nil
}

// SearchResponse represents the response structure for search operations
type SearchResponse struct {
	Result []SearchResult `json:"result"`
}

// SearchResult represents an individual search result item
type SearchResult struct {
	Href string `json:"href"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// GenericCRUDConfig holds configuration for generic CRUD operations
type GenericCRUDConfig struct {
	BaseURL                string
	APIVersion             string
	Auth                   *auth.Auth
	ResourceType           string // e.g., "host", "network", etc.
	GetJSONHeaders         func() map[string]string
	GetAuthHeaders         func() map[string]string
	GetJSONHeadersWithEtag func(string) map[string]string
	GetAuthHeadersWithEtag func(string) map[string]string
}

// CreateResource creates a new resource in the SMC system using generic CRUD operations
// todo pass directly json string payload instead of interface{}
func CreateResource(config *GenericCRUDConfig, resource interface{}) (*ResponseData, error) {
	reqBody, err := json.Marshal(resource)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resource: %w", err)
	}

	createURL := fmt.Sprintf("%s/%s/elements/%s", config.BaseURL, config.APIVersion, config.ResourceType)
	headers := config.GetJSONHeaders()

	resp, err := DoRequest(Options{
		Method:  "POST",
		URL:     createURL,
		Headers: headers,
		Body:    reqBody,
		Timeout: 30 * time.Second,
		Auth:    config.Auth,
	})
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("create %s failed with status %d: %s", config.ResourceType, resp.StatusCode, string(resp.Body))
	}

	return resp, nil
}

// CreateSubResource creates a new resource in the SMC system using generic CRUD operations
// todo pass directly json string payload instead of interface{}
func CreateSubResource(config *GenericCRUDConfig, resource interface{}, url string) (*ResponseData, error) {
	reqBody, err := json.Marshal(resource)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resource: %w", err)
	}

	headers := config.GetJSONHeaders()

	resp, err := DoRequest(Options{
		Method:  "POST",
		URL:     url,
		Headers: headers,
		Body:    reqBody,
		Timeout: 30 * time.Second,
		Auth:    config.Auth,
	})
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("create %s failed with status %d: %s", config.ResourceType, resp.StatusCode, string(resp.Body))
	}

	return resp, nil
}

// ReadResourceByName retrieves a resource by name from the SMC system using generic CRUD operations
func ReadResourceByName(config *GenericCRUDConfig, name string) (*ResponseData, error) {
	encodedName := url.QueryEscape(name)
	searchURL := fmt.Sprintf("%s/%s/elements/%s?filter=%s", config.BaseURL, config.APIVersion, config.ResourceType, encodedName)
	headers := config.GetJSONHeaders()
	searchResp, err := DoRequest(Options{
		Method:  "GET",
		URL:     searchURL,
		Headers: headers,
		Timeout: 30 * time.Second,
		Auth:    config.Auth,
	})
	if err != nil {
		return nil, fmt.Errorf("search request failed: %w", err)
	}

	if searchResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search failed with status %d: %s", searchResp.StatusCode, string(searchResp.Body))
	}

	var searchResponse SearchResponse
	if err := json.Unmarshal(searchResp.Body, &searchResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal search response: %w", err)
	}

	if len(searchResponse.Result) == 0 {
		return nil, fmt.Errorf("no %s found with name: %s", config.ResourceType, name)
	}

	href := searchResponse.Result[0].Href
	resp, err := DoRequest(Options{
		Method:  "GET",
		URL:     href,
		Headers: config.GetAuthHeaders(),
		Timeout: 30 * time.Second,
		Auth:    config.Auth,
	})
	if err != nil {
		return nil, fmt.Errorf("read request failed: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("read %s failed with status %d: %s", config.ResourceType, resp.StatusCode, string(resp.Body))
	}

	// Add href and etag to response for caller convenience
	resp.Location = href
	return resp, nil
}

// ReadResourceByHref retrieves a resource by href from the SMC system
func ReadResourceByHref(config *GenericCRUDConfig, href string) (*ResponseData, error) {
	resp, err := DoRequest(Options{
		Method:  "GET",
		URL:     href,
		Headers: config.GetAuthHeaders(),
		Timeout: 30 * time.Second,
		Auth:    config.Auth,
	})
	if err != nil {
		return nil, fmt.Errorf("read request failed: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("read %s failed with status %d: %s", config.ResourceType, resp.StatusCode, string(resp.Body))
	}

	return resp, nil
}

// UpdateResource updates an existing resource in the SMC system using generic CRUD operations
func UpdateResource(config *GenericCRUDConfig, resource interface{}, resourceName string) (*ResponseData, error) {
	// First, read the current resource to get ETag and Href
	currentResp, err := ReadResourceByName(config, resourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to read current %s for update: %w", config.ResourceType, err)
	}
	if currentResp == nil {
		return nil, fmt.Errorf("%s not found for update: %s", config.ResourceType, resourceName)
	}

	reqBody, err := json.Marshal(resource)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resource: %w", err)
	}
	fmt.Printf("UpdateResource request body: %s\n", string(reqBody))
	headers := config.GetJSONHeadersWithEtag(currentResp.ETag)
	resp, err := DoRequest(Options{
		Method:  "PUT",
		URL:     currentResp.Location,
		Headers: headers,
		Body:    reqBody,
		Timeout: 30 * time.Second,
		Auth:    config.Auth,
	})
	if err != nil {
		return nil, fmt.Errorf("update request failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("update %s failed with status %d: %s", config.ResourceType, resp.StatusCode, string(resp.Body))
	}

	return resp, nil
}

// DeleteResource deletes a resource from the SMC system by name using generic CRUD operations
func DeleteResource(config *GenericCRUDConfig, resourceName string) error {
	// First, read the resource to get ETag and Href
	resourceResp, err := ReadResourceByName(config, resourceName)
	if err != nil {
		return fmt.Errorf("failed to read %s for deletion: %w", config.ResourceType, err)
	}
	if resourceResp == nil {
		return fmt.Errorf("%s not found for deletion: %s", config.ResourceType, resourceName)
	}

	headers := config.GetAuthHeadersWithEtag(resourceResp.ETag)
	resp, err := DoRequest(Options{
		Method:  "DELETE",
		URL:     resourceResp.Location,
		Headers: headers,
		Timeout: 30 * time.Second,
		Auth:    config.Auth,
	})
	if err != nil {
		return fmt.Errorf("delete request failed: %w", err)
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("delete %s failed with status %d: %s", config.ResourceType, resp.StatusCode, string(resp.Body))
	}

	return nil
}

// DeleteResourceByHref deletes a resource from the SMC system by href using generic CRUD operations
func DeleteResourceByHref(config *GenericCRUDConfig, href, etag string) error {
	headers := config.GetAuthHeadersWithEtag(etag)
	resp, err := DoRequest(Options{
		Method:  "DELETE",
		URL:     href,
		Headers: headers,
		Timeout: 30 * time.Second,
		Auth:    config.Auth,
	})
	if err != nil {
		return fmt.Errorf("delete request failed: %w", err)
	}

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("delete %s failed with status %d: %s", config.ResourceType, resp.StatusCode, string(resp.Body))
	}

	return nil
}

// SearchElements performs a search and returns all matching elements without fetching their details
func SearchElements(config *GenericCRUDConfig, namePattern string) ([]SearchResult, error) {
	encodedName := url.QueryEscape(namePattern)
	searchURL := fmt.Sprintf("%s/%s/elements/%s?filter=%s&exact_match=true", config.BaseURL, config.APIVersion, config.ResourceType, encodedName)
	headers := config.GetJSONHeaders()

	searchResp, err := DoRequest(Options{
		Method:  "GET",
		URL:     searchURL,
		Headers: headers,
		Timeout: 30 * time.Second,
		Auth:    config.Auth,
	})
	if err != nil {
		return nil, fmt.Errorf("search request failed: %w", err)
	}

	if searchResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search failed with status %d: %s", searchResp.StatusCode, string(searchResp.Body))
	}

	var searchResponse SearchResponse
	if err := json.Unmarshal(searchResp.Body, &searchResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal search response: %w", err)
	}

	return searchResponse.Result, nil
}
