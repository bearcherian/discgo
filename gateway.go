package discgo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// EventHandler is a function that takes a single interface{} argument which
// is an Event Data type as listed in events.go and documented in
// https://discord.com/developers/docs/topics/gateway#commands-and-events-gateway-events
// this is the expected argument for all event handlers, and is intentionally left as an
// interface{} type, instead of multiple handlers with concrete types. This prevents the
// need for a new handler each event message, and allows consumers of the library to use
// their own structs if desired. In order to reduce some code when converting to the desired
//  struct, the function DataToStruct is provided.
type EventHandler func(data interface{})

type Gateway struct {
	dialer *websocket.Dialer
	conn   *websocket.Conn
	logger Logger

	config GatewayConfig

	heartbeatInterval int
	sessionID         string
	lastAck           time.Time
	stopChan          chan int
	lastSequence      int

	eventHandlers map[string][]EventHandler
}

type GatewayConfig struct {
	GatewayURL  string
	AuthToken   string
	ShardID     int
	TotalShards int
}

type GatewayOption func(g *Gateway)

func WithConfig(config GatewayConfig) GatewayOption {
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

func WithLogger(l Logger) GatewayOption {
	return func(g *Gateway) {
		g.logger = l
	}
}

func NewGateway(options ...GatewayOption) *Gateway {
	gateway := Gateway{
		sessionID:     "",
		lastAck:       time.Now(),
		lastSequence:  0,
		stopChan:      make(chan int, 1),
		eventHandlers: make(map[string][]EventHandler),
	}

	for _, opt := range options {
		opt(&gateway)
	}

	if gateway.config.TotalShards == 0 {
		gateway.config.TotalShards = 1
	}

	if gateway.dialer == nil {
		gateway.dialer = websocket.DefaultDialer
	}

	if gateway.logger == nil {
		gateway.logger = &noopLogger{}
	}

	return &gateway
}

//Start establishes and maintains websocket connection with the Discord Gateway
func (g *Gateway) Start() {
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))

	// dial and establish a websocket connection
	conn, _, err := g.dialer.DialContext(ctx, fmt.Sprintf("%s?encoding=json", g.config.GatewayURL), nil)
	if err != nil {
		g.logger.Errorf("failed to establish a gateway connection %w", err)
		return
	}
	g.conn = conn

	// get and read in the next payload, which should be Hello
	var helloPayload Payload
	if err = conn.ReadJSON(&helloPayload); err != nil {
		g.logger.Errorf("unable to parse payload: %w", err)
		return
	}
	if helloPayload.Op != OpCodeHello {
		g.logger.Errorf("expected op code %d but got %d", OpCodeHello, helloPayload.Op)
		return
	}

	var helloData Hello
	if err := helloPayload.MarshalData(&helloData); err != nil {
		g.logger.Error("cannot read hello data: ", err)
	}

	// send identify
	if err := g.sendIdentify(); err != nil {
		g.logger.Errorf("Identify failed: %v", err)
		return
	}

	// if we have an existing session ID, then we need to resume
	if g.sessionID != "" {
		if err := g.resumeConnection(); err != nil {
			g.logger.Errorf("failed to resume connection: %v", err)
			return
		}
	}

	// start heartbeating
	go g.startHeartBeat(helloData.HeartbeatInterval)

	go g.startListening()
}

// AddEventHandlers adds and registers a function that will be called when the bot recieves
// a message for the specified event type. See EventHandler for the expected func, and
// DataToStruct on how to handle the function argument
func (g *Gateway) AddEventHandler(eventType string, handler EventHandler) {
	g.eventHandlers[eventType] = append(g.eventHandlers[eventType], handler)
}

func (g *Gateway) sendHeartBeat() {
	g.conn.WriteJSON(Payload{
		Op: OpCodeHeartbeat,
		D:  g.lastSequence,
	})
}

func (g *Gateway) sendIdentify() error {
	intents := IntentMessageCreate +
		IntentMessageUpdate +
		IntentMessageDelete

	if err := g.conn.WriteJSON(Payload{
		Op: OpCodeIdentify,
		D: Identify{
			Token:   g.config.AuthToken,
			Intents: intents,
			Properties: IdentifyConnectionProperties{
				OS:      "docker",
				Browser: "teamtool-server",
				Device:  "teamtool-server",
			},
			Shard: []int{0, 1},
		},
	}); err != nil {
		return fmt.Errorf("unable to send Identify message: %w", err)
	}

	g.logger.Info("bot connected")
	return nil
}

func (g *Gateway) startHeartBeat(interval int) {

	g.logger.Infof("starting heartbeat with interval of %d", interval)

	t := time.NewTicker(time.Duration(interval) * time.Millisecond)

	for {
		select {
		case <-g.stopChan:
			g.logger.Info("stopping heartbeats")
			return
		case <-t.C:
			nowMs := time.Now().Nanosecond() / 1000 / 1000

			g.logger.Debug("sending heartbeat...")
			// if we haven't recieved a heartbeat since the last interval, we need to reconnect
			if nowMs-g.lastAck.Nanosecond()/1000/1000 > int(interval) {
				g.CloseWithCode(websocket.CloseServiceRestart)
				g.conn.Close()
				g.logger.Info("no ACK since last heartbeat, reconecting")

				// restart the connection
				go g.Start()
			}
			g.sendHeartBeat()
		}
	}
}

func (g *Gateway) startListening() {
	defer g.logger.Info("done reading")

	for {

		if mType, reader, err := g.conn.NextReader(); err != nil {
			if closeErr, ok := err.(*websocket.CloseError); ok {
				g.logger.Info("connection closed: ", closeErr)
				return
			}
			g.logger.Error("reader failed: ", err)
			g.Close()
			return
		} else {

			g.logger.Debugf("New Reader message of type %s", mType)

			switch mType {
			case websocket.TextMessage:
				g.handleTextMessage(reader)
			case websocket.BinaryMessage:
				g.logger.Info("unhandled binary message")
			case websocket.CloseMessage:
				g.logger.Info("unhandled close message")
			case websocket.PingMessage:
				g.logger.Info("unhandled ping message")
			case websocket.PongMessage:
				g.logger.Info("unhandled pong message")
			default:
				g.logger.Error("unknown message type: ", mType)
			}
		}
	}
}

func (g *Gateway) resumeConnection() error {
	return g.conn.WriteJSON(Payload{
		Op: OpCodeResume,
		D: Resume{
			Token:     g.config.AuthToken,
			SessionId: g.sessionID,
			Seq:       g.lastSequence,
		},
	})

}

func (g *Gateway) CloseWithCode(code int) error {
	return g.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(code, ""))
}

func (g *Gateway) Close() error {
	g.logger.Info("closing connection")
	g.stopChan <- 1
	return g.CloseWithCode(websocket.CloseGoingAway)
}

// DataToStruct takes a JSON marshallable struct and Unmarshals it to the target pointer value
func DataToStruct(source interface{}, target interface{}) error {
	bytes, err := json.Marshal(source)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, target)
}

const (
	GetGatewayBotURL = "/gateway/bot"
)

type GetGatewayBotResponse struct {
	// The WSS URL that can be used for connecting to the gateway
	Url string `json:"url"`
	// The recommended number of shards to use when connecting
	Shards int `json:"shards"`
	// Information on the current session start limit
	SessionStartLimit SessionStartLimit `json:"session_start_limit"`
}

func (d *ApiClient) GetGatewayBot(ctx context.Context) (*GetGatewayBotResponse, *http.Response, error) {

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s?encoding=json", d.discordConfig.APIBaseURL, GetGatewayBotURL), nil)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to build request: %w", err)
	}

	resp, err := d.Do(request)
	if err != nil {
		return nil, resp, fmt.Errorf("unable to complete request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		errBytes, _ := ioutil.ReadAll(resp.Body)
		return nil, resp, fmt.Errorf("reqeust failed: %d - %s: %s", resp.StatusCode, resp.Status, string(errBytes))
	}

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, fmt.Errorf("unable to read response body: %w", err)
	}

	var response GetGatewayBotResponse
	if err = json.Unmarshal(responseBytes, &response); err != nil {
		return nil, resp, fmt.Errorf("unable to parse JSON: %w", err)
	}

	return &response, resp, nil
}
