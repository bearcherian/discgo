package client

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Config struct {
	APIBaseURL    string
	ApplicationID string
	PublicKey     string
	ClientID      string
	ClientSecret  string
	BotToken      string
}
type Client struct {
	discordConfig *Config
	client        HTTPClient
	logger        *zap.SugaredLogger
}

func NewClient(options ...DiscordClientOption) *Client {
	d := Client{}

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
		d.logger = zap.NewNop().Sugar()
	}

	return &d
}

// Do adds headers before calling DiscordClient.http.Do
func (d *Client) Do(req *http.Request) (*http.Response, error) {

	req.Header.Set("Authorization", "Bot "+d.discordConfig.BotToken)

	fmt.Printf("request: \n\t%#v\n", req)

	return d.client.Do(req)
}

func (d *Client) Close() {
	// nothing to do here right now
}
