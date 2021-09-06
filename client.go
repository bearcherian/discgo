package discgo

import (
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type ApiConfig struct {
	APIBaseURL    string
	ApplicationID string
	PublicKey     string
	ClientID      string
	ClientSecret  string
	BotToken      string
}
type ApiClient struct {
	discordConfig *ApiConfig
	client        HTTPClient
	logger        Logger
}

type ApiClientOption func(*ApiClient)

func WithApiConfig(cfg *ApiConfig) ApiClientOption {
	return func(d *ApiClient) {
		d.discordConfig = cfg
	}
}

func WithHttpClient(c HTTPClient) ApiClientOption {
	return func(d *ApiClient) {
		d.client = c
	}
}

func WithClientLogger(l Logger) ApiClientOption {
	return func(dc *ApiClient) {
		// TODO make this a logger interface
		dc.logger = l
	}
}

func NewApiClient(options ...ApiClientOption) *ApiClient {
	d := ApiClient{}

	for _, opt := range options {
		opt(&d)
	}

	if d.client == nil {
		d.client = http.DefaultClient
	}

	if d.discordConfig.APIBaseURL == "" {
		d.discordConfig.APIBaseURL = "https://discord.com/api"
	}

	if d.logger == nil {
		d.logger = &noopLogger{}
	}

	return &d
}

// Do adds headers before calling DiscordClient.http.Do
func (d *ApiClient) Do(req *http.Request) (*http.Response, error) {

	req.Header.Set("Authorization", "Bot "+d.discordConfig.BotToken)

	return d.client.Do(req)
}

func (d *ApiClient) Close() {
	// nothing to do here right now
}
