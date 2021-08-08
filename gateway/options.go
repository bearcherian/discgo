package gateway

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type GatewayOption func(g *Gateway)

type Config struct {
	GatewayURL  string
	AuthToken   string
	ShardID     int
	TotalShards int
}

func WithConfig(config Config) GatewayOption {
	return func(g *Gateway) {
		g.config = config
	}
}

// WithDialer specifies a custom websocket.Dialer to use when connecting to the Discord Gateway
func WithDialer(d *websocket.Dialer) GatewayOption {
	return func(g *Gateway) {
		g.dialer = d
	}
}

func WithLogger(l *zap.SugaredLogger) GatewayOption {
	return func(g *Gateway) {
		g.logger = l
	}
}
