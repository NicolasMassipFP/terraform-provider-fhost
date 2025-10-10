package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func TestNewAuth(t *testing.T) {
	url := "https://example.com"
	apiToken := "token123"
	apiKey := "key456"
	apiVersion := "v1"
	insecure := true
	trustedCert := "cert-data"

	auth := NewAuth(url, apiToken, apiKey, apiVersion, &insecure, trustedCert)

	if auth.URL != url {
		t.Errorf("Expected URL %s, got %s", url, auth.URL)
	}
	if auth.APIToken != apiToken {
		t.Errorf("Expected APIToken %s, got %s", apiToken, auth.APIToken)
	}
	if auth.APIKey != apiKey {
		t.Errorf("Expected APIKey %s, got %s", apiKey, auth.APIKey)
	}
	if auth.APIVersion != apiVersion {
		t.Errorf("Expected APIVersion %s, got %s", apiVersion, auth.APIVersion)
	}
	if auth.Insecure == nil || *auth.Insecure != insecure {
		t.Errorf("Expected Insecure %v, got %v", insecure, auth.Insecure)
	}
	if auth.TrustedCert != trustedCert {
		t.Errorf("Expected TrustedCert %s, got %s", trustedCert, auth.TrustedCert)
	}
}

func TestNewAuth_NilInsecure(t *testing.T) {
	auth := NewAuth("url", "token", "key", "v1", nil, "cert")
	if auth.Insecure != nil {
		t.Errorf("Expected Insecure to be nil, got %v", auth.Insecure)
	}
}

func TestGetEnvURL(t *testing.T) {
	// Test when environment variable is set
	expectedURL := "https://smc.example.com"
	_ = os.Setenv("SMC_URL", expectedURL)
	defer func() { _ = os.Unsetenv("SMC_URL") }()

	auth := &Auth{}
	url, err := auth.GetEnvURL()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if url != expectedURL {
		t.Errorf("Expected URL %s, got %s", expectedURL, url)
	}
	if auth.URL != expectedURL {
		t.Errorf("Expected auth.URL %s, got %s", expectedURL, auth.URL)
	}
}

func TestGetEnvURL_NotSet(t *testing.T) {
	// Ensure environment variable is not set
	_ = os.Unsetenv("SMC_URL")

	auth := &Auth{}
	url, err := auth.GetEnvURL()

	if err == nil {
		t.Error("Expected error when SMC_URL is not set, got nil")
	}
	if url != "" {
		t.Errorf("Expected empty URL, got %s", url)
	}
	expectedError := "SMC_URL environment variable not set"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %s, got %s", expectedError, err.Error())
	}
}

func TestGetEnvAPIKey(t *testing.T) {
	// Test when environment variable is set
	expectedKey := "test-api-key"
	_ = os.Setenv("SMC_API_KEY", expectedKey)
	defer func() { _ = os.Unsetenv("SMC_API_KEY") }()

	auth := &Auth{}
	key, err := auth.GetEnvAPIKey()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if key != expectedKey {
		t.Errorf("Expected API key %s, got %s", expectedKey, key)
	}
	if auth.APIKey != expectedKey {
		t.Errorf("Expected auth.APIKey %s, got %s", expectedKey, auth.APIKey)
	}
}

func TestGetEnvAPIKey_NotSet(t *testing.T) {
	// Ensure environment variable is not set
	_ = os.Unsetenv("SMC_API_KEY")

	auth := &Auth{}
	key, err := auth.GetEnvAPIKey()

	if err == nil {
		t.Error("Expected error when SMC_API_KEY is not set, got nil")
	}
	if key != "" {
		t.Errorf("Expected empty key, got %s", key)
	}
	expectedError := "SMC_API_KEY environment variable not set"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %s, got %s", expectedError, err.Error())
	}
}

func TestGetEnvAPIToken(t *testing.T) {
	// Test when environment variable is set
	expectedToken := "test-api-token"
	_ = os.Setenv("SMC_API_TOKEN", expectedToken)
	defer func() { _ = os.Unsetenv("SMC_API_TOKEN") }()

	auth := &Auth{}
	token, err := auth.GetEnvAPIToken()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if token != expectedToken {
		t.Errorf("Expected API token %s, got %s", expectedToken, token)
	}
	if auth.APIToken != expectedToken {
		t.Errorf("Expected auth.APIToken %s, got %s", expectedToken, auth.APIToken)
	}
}

func TestGetEnvAPIToken_NotSet(t *testing.T) {
	// Ensure environment variable is not set
	_ = os.Unsetenv("SMC_API_TOKEN")

	auth := &Auth{}
	token, err := auth.GetEnvAPIToken()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if token != "" {
		t.Errorf("Expected empty token, got %s", token)
	}
	if auth.APIToken != "" {
		t.Errorf("Expected auth.APIToken to be empty, got %s", auth.APIToken)
	}
}

func TestGetEnvAPIVersion(t *testing.T) {
	// Test when environment variable is set
	expectedVersion := "v2"
	_ = os.Setenv("SMC_API_VERSION", expectedVersion)
	defer func() { _ = os.Unsetenv("SMC_API_VERSION") }()

	auth := &Auth{}
	version, err := auth.GetEnvAPIVersion()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if version != expectedVersion {
		t.Errorf("Expected API version %s, got %s", expectedVersion, version)
	}
	if auth.APIVersion != expectedVersion {
		t.Errorf("Expected auth.APIVersion %s, got %s", expectedVersion, auth.APIVersion)
	}
}

func TestGetEnvAPIVersion_NotSet(t *testing.T) {
	// Ensure environment variable is not set
	_ = os.Unsetenv("SMC_API_VERSION")

	auth := &Auth{}
	version, err := auth.GetEnvAPIVersion()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if version != "" {
		t.Errorf("Expected empty version, got %s", version)
	}
	if auth.APIVersion != "" {
		t.Errorf("Expected auth.APIVersion to be empty, got %s", auth.APIVersion)
	}
}

func TestGetEnvInsecure(t *testing.T) {
	// Test when environment variable is set to "true"
	_ = os.Setenv("SMC_INSECURE", "true")
	defer func() { _ = os.Unsetenv("SMC_INSECURE") }()

	auth := &Auth{}
	insecure, err := auth.GetEnvInsecure()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !insecure {
		t.Error("Expected insecure to be true")
	}
	if auth.Insecure == nil || !*auth.Insecure {
		t.Error("Expected auth.Insecure to be true")
	}
}

func TestGetEnvInsecure_False(t *testing.T) {
	// Test when environment variable is set to something other than "true"
	_ = os.Setenv("SMC_INSECURE", "false")
	defer func() { _ = os.Unsetenv("SMC_INSECURE") }()

	auth := &Auth{}
	insecure, err := auth.GetEnvInsecure()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if insecure {
		t.Error("Expected insecure to be false")
	}
	if auth.Insecure == nil || *auth.Insecure {
		t.Error("Expected auth.Insecure to be false")
	}
}

func TestGetEnvInsecure_NotSet(t *testing.T) {
	// Test when environment variable is not set
	_ = os.Unsetenv("SMC_INSECURE")

	auth := &Auth{}
	insecure, err := auth.GetEnvInsecure()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if insecure {
		t.Error("Expected insecure to be false when not set")
	}
	if auth.Insecure == nil || *auth.Insecure {
		t.Error("Expected auth.Insecure to be false when not set")
	}
}

// Test combination scenarios
func TestAuth_AllEnvVariables(t *testing.T) {
	// Set all environment variables
	_ = os.Setenv("SMC_URL", "https://test.com")
	_ = os.Setenv("SMC_API_KEY", "test-key")
	_ = os.Setenv("SMC_API_TOKEN", "test-token")
	_ = os.Setenv("SMC_API_VERSION", "v3")
	_ = os.Setenv("SMC_INSECURE", "true")

	defer func() {
		_ = os.Unsetenv("SMC_URL")
		_ = os.Unsetenv("SMC_API_KEY")
		_ = os.Unsetenv("SMC_API_TOKEN")
		_ = os.Unsetenv("SMC_API_VERSION")
		_ = os.Unsetenv("SMC_INSECURE")
	}()

	auth := &Auth{}

	// Test each method
	url, _ := auth.GetEnvURL()
	key, _ := auth.GetEnvAPIKey()
	token, _ := auth.GetEnvAPIToken()
	version, _ := auth.GetEnvAPIVersion()
	insecure, _ := auth.GetEnvInsecure()

	if url != "https://test.com" {
		t.Errorf("Expected URL https://test.com, got %s", url)
	}
	if key != "test-key" {
		t.Errorf("Expected key test-key, got %s", key)
	}
	if token != "test-token" {
		t.Errorf("Expected token test-token, got %s", token)
	}
	if version != "v3" {
		t.Errorf("Expected version v3, got %s", version)
	}
	if !insecure {
		t.Error("Expected insecure to be true")
	}

	// Verify auth struct is updated
	if auth.URL != "https://test.com" ||
		auth.APIKey != "test-key" ||
		auth.APIToken != "test-token" ||
		auth.APIVersion != "v3" ||
		auth.Insecure == nil ||
		!*auth.Insecure {
		t.Error("Auth struct was not properly updated")
	}
}

// Helper function to generate a test certificate for auth tests
func generateTestCertificateForAuth() (string, error) {
	// Generate private key
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}

	// Create certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "localhost",
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses: []net.IP{net.IPv4(127, 0, 0, 1)},
		DNSNames:    []string{"localhost"},
	}

	// Create certificate
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return "", err
	}

	// Create PEM encoded certificate
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	return string(certPEM), nil
}

// TestNewAuth_WithTrustedCert tests auth creation with a trusted certificate
func TestNewAuth_WithTrustedCert(t *testing.T) {
	certPEM, err := generateTestCertificateForAuth()
	if err != nil {
		t.Fatalf("Failed to generate test certificate: %v", err)
	}

	auth := NewAuth("https://example.com", "token", "key", "v1", nil, certPEM)

	if auth.TrustedCert != certPEM {
		t.Errorf("Expected TrustedCert to be set correctly")
	}

	// Verify certificate can be parsed
	block, _ := pem.Decode([]byte(auth.TrustedCert))
	if block == nil {
		t.Error("Expected valid PEM certificate block")
		return
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Errorf("Expected valid certificate, got error: %v", err)
		return
	}

	if cert.Subject.CommonName != "localhost" {
		t.Errorf("Expected CommonName 'localhost', got %s", cert.Subject.CommonName)
	}
}

// TestAuth_HTTPSModeDetection tests if auth detects HTTPS mode based on TrustedCert
func TestAuth_HTTPSModeDetection(t *testing.T) {
	tests := []struct {
		name        string
		trustedCert string
		expectHTTPS bool
	}{
		{
			name:        "Empty certificate - HTTP mode",
			trustedCert: "",
			expectHTTPS: false,
		},
		{
			name:        "Valid certificate - HTTPS mode",
			trustedCert: "dummy-cert-content",
			expectHTTPS: true,
		},
		{
			name:        "Whitespace certificate - HTTP mode",
			trustedCert: "   \n\t  ",
			expectHTTPS: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := NewAuth("http://example.com", "token", "key", "v1", nil, tt.trustedCert)

			// Check if HTTPS mode should be enabled
			httpsMode := strings.TrimSpace(auth.TrustedCert) != ""
			if httpsMode != tt.expectHTTPS {
				t.Errorf("Expected HTTPS mode %v, got %v", tt.expectHTTPS, httpsMode)
			}
		})
	}
}

// TestAuth_InsecureWithTrustedCert tests behavior when both insecure and trusted cert are set
func TestAuth_InsecureWithTrustedCert(t *testing.T) {
	certPEM, err := generateTestCertificateForAuth()
	if err != nil {
		t.Fatalf("Failed to generate test certificate: %v", err)
	}

	insecure := true
	auth := NewAuth("https://example.com", "token", "key", "v1", &insecure, certPEM)

	// Both insecure and trusted cert are set - insecure should take precedence
	if auth.Insecure == nil || !*auth.Insecure {
		t.Error("Expected insecure mode to be enabled")
	}

	if auth.TrustedCert != certPEM {
		t.Error("Expected trusted cert to be preserved even when insecure is set")
	}
}
