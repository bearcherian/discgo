package client

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type DiscordConfig struct {
	APIBaseURL    string
	ApplicationID string
	PublicKey     string
	ClientID      string
	ClientSecret  string
	BotToken      string
}
type DiscordClient struct {
	discordConfig *DiscordConfig
	client        HTTPClient
	logger        *zap.SugaredLogger

	// websocket dialer and gateway connection
	// dialer  *websocket.Dialer
	// gateway *websocket.Conn
}

func NewClient(options ...DiscordClientOption) *DiscordClient {
	d := DiscordClient{}

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
func (d *DiscordClient) Do(req *http.Request) (*http.Response, error) {

	req.Header.Set("Authorization", "Bot "+d.discordConfig.BotToken)

	fmt.Printf("request: \n\t%#v\n", req)

	return d.client.Do(req)
}

func (d *DiscordClient) Close() {
	// nothing to do here right now
}
