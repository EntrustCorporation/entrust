package entrust

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
)

// Default client configurations
const (
	APIServer = "https://api.entrust.net/enterprise/v2"
)

// Client for the Entrust API
type Client struct {
	username string
	apiKey   string
	client   *http.Client
}

// New Entrust API client
func New() (*Client, error) {
	var err error
	var clientCertificate tls.Certificate

	certPEM := os.Getenv("ENTRUST_API_CERTIFICATE")
	certPath := os.Getenv("ENTRUST_API_CERTIFICATE_PATH")

	if certPEM != "" {
		clientCertificate, err = tls.X509KeyPair([]byte(certPEM), []byte(os.Getenv("ENTRUST_API_PRIVATE_KEY")))
	} else if certPath != "" {
		clientCertificate, err = tls.LoadX509KeyPair(certPath, os.Getenv("ENTRUST_API_PRIVATE_KEY_PATH"))
	} else {
		return nil, fmt.Errorf("invalid configuration, no client certificate provided")
	}
	if err != nil {
		return nil, fmt.Errorf("error loading client certificate key pair: %w", err)
	}

	return &Client{
		username: os.Getenv("ENTRUST_API_USERNAME"),
		apiKey:   os.Getenv("ENTRUST_API_PASSEWORD"),
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					Certificates: []tls.Certificate{clientCertificate},
				},
			},
		},
	}, nil
}
