// Package auth provides authentication utilities for the SMC provider.
package auth

import (
	"fmt"
	"os"
)

// Auth describes the authentication information for SMC
type Auth struct {
	URL         string
	APIToken    string
	APIKey      string
	APIVersion  string
	Insecure    *bool
	TrustedCert string
}

// NewAuth initializes Auth object with the given metadata
func NewAuth(url, apiToken, apiKey, apiVersion string, insecure *bool, trustedCert string) *Auth {
	return &Auth{
		URL:         url,
		APIToken:    apiToken,
		APIKey:      apiKey,
		APIVersion:  apiVersion,
		Insecure:    insecure,
		TrustedCert: trustedCert,
	}
}

// GetEnvURL gets SMC URL from OS environment
func (a *Auth) GetEnvURL() (string, error) {
	url := os.Getenv("SMC_URL")
	if url == "" {
		return url, fmt.Errorf("SMC_URL environment variable not set")
	}
	a.URL = url
	return url, nil
}

// GetEnvAPIKey gets API key from OS environment
func (a *Auth) GetEnvAPIKey() (string, error) {
	key := os.Getenv("SMC_API_KEY")
	if key == "" {
		return key, fmt.Errorf("SMC_API_KEY environment variable not set")
	}
	a.APIKey = key
	return key, nil
}

// GetEnvAPIToken gets API token from OS environment
func (a *Auth) GetEnvAPIToken() (string, error) {
	token := os.Getenv("SMC_API_TOKEN")
	a.APIToken = token
	return token, nil
}

// GetEnvAPIVersion gets API version from OS environment
func (a *Auth) GetEnvAPIVersion() (string, error) {
	v := os.Getenv("SMC_API_VERSION")
	a.APIVersion = v
	return v, nil
}

// GetEnvInsecure gets Insecure value from OS environment
func (a *Auth) GetEnvInsecure() (bool, error) {
	c := os.Getenv("SMC_INSECURE")
	if c == "true" {
		b := true
		a.Insecure = &b
		return true, nil
	}
	b := false
	a.Insecure = &b
	return false, nil
}
