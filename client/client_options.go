package client

import (
	"go.uber.org/zap"
)

type DiscordClientOption func(*Client)

func WithDiscordConfig(cfg *Config) DiscordClientOption {
	return func(d *Client) {
		d.discordConfig = cfg
	}
}

func WithHttpClient(c HTTPClient) DiscordClientOption {
	return func(d *Client) {
		d.client = c
	}
}

func WithLogger(l *zap.SugaredLogger) DiscordClientOption {
	return func(dc *Client) {
		dc.logger = l
	}
}
