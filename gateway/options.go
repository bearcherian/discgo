package gateway

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type GatewayOption func(g *Gateway)

func WithAuthToken(token string) GatewayOption {
	return func(g *Gateway) {
		g.authToken = token
	}
}

func WithGatewayURL(gatewayUrl string) GatewayOption {
	return func(g *Gateway) {
		g.gatewayURL = gatewayUrl
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
