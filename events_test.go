package discgo

import (
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func TestGateway_handleEventHello(t *testing.T) {
	type fields struct {
		dialer            *websocket.Dialer
		conn              *websocket.Conn
		logger            *zap.SugaredLogger
		authToken         string
		gatewayURL        string
		heartbeatInterval int
		sessionID         string
		lastAck           time.Time
		stopChan          chan int
		lastSequence      int
		eventHandlers     map[string][]EventHandler
	}
	type args struct {
		p Payload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gateway{
				dialer: tt.fields.dialer,
				conn:   tt.fields.conn,
				logger: tt.fields.logger,
				config: Config{
					AuthToken:  tt.fields.authToken,
					GatewayURL: tt.fields.gatewayURL,
				},
				heartbeatInterval: tt.fields.heartbeatInterval,
				sessionID:         tt.fields.sessionID,
				lastAck:           tt.fields.lastAck,
				stopChan:          tt.fields.stopChan,
				lastSequence:      tt.fields.lastSequence,
				eventHandlers:     tt.fields.eventHandlers,
			}
			if err := g.handleEventHello(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Gateway.handleEventHello() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
