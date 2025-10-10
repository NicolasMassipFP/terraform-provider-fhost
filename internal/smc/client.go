// Package smc provides the core functionality for the SMC Terraform provider.
package smc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	sdk_auth "github.com/terraform-providers/terraform-provider-smc/internal/sdk/auth"
	sdk_request "github.com/terraform-providers/terraform-provider-smc/internal/sdk/request"
)

// Client represents an SMC API client with authentication capabilities.
type Client struct {
	URL           string
	APIKey        string
	VerifySSL     bool
	Token         string
	HTTPClient    *http.Client
	APIVersion    string
	Auth          *sdk_auth.Auth
	UseAuthHeader bool // Track whether to use Authorization header or Cookie
}

// LoginRequest represents the request payload for SMC authentication.
type LoginRequest struct {
	APIKey string `json:"authenticationkey"`
}

/* type LoginResponse struct {
	SessionID string `json:"session_id"`
} */

// NewClientFromAuth creates a new Client from an Auth struct
func NewClientFromAuth(auth *sdk_auth.Auth) (*Client, error) {

	tlsConfig := &tls.Config{InsecureSkipVerify: auth.Insecure != nil && *auth.Insecure}
	if auth.TrustedCert != "" {
		// Load custom trusted cert
		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM([]byte(auth.TrustedCert)) {
			return nil, fmt.Errorf("failed to append trusted certificate")
		}
		tlsConfig.RootCAs = certPool
	}
	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	httpClient := &http.Client{Transport: tr}

	c := &Client{
		URL:        auth.URL,
		APIKey:     auth.APIKey,
		VerifySSL:  auth.Insecure == nil || !*auth.Insecure,
		HTTPClient: httpClient,
		APIVersion: auth.APIVersion,
		Auth:       auth,
	}

	err := c.Login(context.Background())
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Login authenticates with the SMC API using the provided API key.
func (c *Client) Login(ctx context.Context) error {
	loginRequest := LoginRequest{
		APIKey: c.APIKey,
	}

	reqBody, err := json.Marshal(loginRequest)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/%s/login", c.URL, c.APIVersion)
	opts := sdk_request.Options{
		Method:  "POST",
		URL:     url,
		Headers: map[string]string{"Content-Type": "application/json"},
		Body:    reqBody,
		Context: ctx,
		Auth:    c.Auth,
	}
	resp, err := sdk_request.DoRequest(opts)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("login failed with status %d: %s", resp.StatusCode, string(resp.Body))
	}

	// Try to get token from Authorization header first (for HTTPS)
	if authHeader := resp.Headers.Get("Authorization"); authHeader != "" {
		c.Token = authHeader
		c.UseAuthHeader = true
		return nil
	}

	// Fallback: Extract JSESSIONID from Set-Cookie header (for HTTP)
	cookies := resp.Headers["Set-Cookie"]
	for _, cookieStr := range cookies {
		if strings.Contains(cookieStr, "JSESSIONID=") {
			parts := strings.Split(cookieStr, ";")
			for _, part := range parts {
				if strings.HasPrefix(part, "JSESSIONID=") {
					c.Token = strings.TrimPrefix(part, "JSESSIONID=")
					c.UseAuthHeader = false
					break
				}
			}
		}
	}
	return nil
}

// GetJSONHeaders returns standard headers for JSON API requests with authentication
func (c *Client) GetJSONHeaders() map[string]string {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	if c.UseAuthHeader {
		headers["Authorization"] = c.Token
	} else {
		headers["Cookie"] = fmt.Sprintf("JSESSIONID=%s", c.Token)
	}

	return headers
}

// GetAuthHeaders returns headers for authenticated requests (no Content-Type)
func (c *Client) GetAuthHeaders() map[string]string {
	if c.UseAuthHeader {
		return map[string]string{
			"Authorization": c.Token,
		}
	}
	return map[string]string{
		"Cookie": fmt.Sprintf("JSESSIONID=%s", c.Token),
	}
}

// GetJSONHeadersWithEtag returns JSON headers with If-Match header for conditional updates
func (c *Client) GetJSONHeadersWithEtag(etag string) map[string]string {
	headers := c.GetJSONHeaders()
	headers["If-Match"] = etag
	return headers
}

// GetAuthHeadersWithEtag returns auth headers with If-Match header for conditional operations
func (c *Client) GetAuthHeadersWithEtag(etag string) map[string]string {
	headers := c.GetAuthHeaders()
	headers["If-Match"] = etag
	return headers
}
