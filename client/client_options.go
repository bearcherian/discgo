package client

import (
	"go.uber.org/zap"
)

type DiscordClientOption func(*DiscordClient)

func WithDiscordConfig(cfg *DiscordConfig) DiscordClientOption {
	return func(d *DiscordClient) {
		d.discordConfig = cfg
	}
}

func WithHttpClient(c HTTPClient) DiscordClientOption {
	return func(d *DiscordClient) {
		d.client = c
	}
}

func WithLogger(l *zap.SugaredLogger) DiscordClientOption {
	return func(dc *DiscordClient) {
		dc.logger = l
	}
}
