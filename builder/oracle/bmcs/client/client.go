package bmcs

import (
	"net/http"
)

const (
	apiVersion     = "20160918"
	userAgent      = "go-bmcs/" + apiVersion
	baseURLPattern = "https://%s.%s.oraclecloud.com/%s/"
)

// Client is the main interface through which consumers interact with the BMCS
// API.
type Client struct {
	UserAgent string
	Compute   *ComputeClient
	Config    *Config
}

// NewClient creates a new Client for communicating with the BMCS API.
func NewClient(config *Config) (*Client, error) {
	transport := NewTransport(http.DefaultTransport, config)
	base := newBaseClient().Client(&http.Client{Transport: transport})

	return &Client{
		UserAgent: userAgent,
		Compute:   NewComputeClient(base.New().Base(config.getBaseURL("iaas"))),
		Config:    config,
	}, nil
}
