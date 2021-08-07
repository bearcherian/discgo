package gateway

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func (g *Gateway) handleTextMessage(reader io.Reader) {

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		g.logger.Errorf("unable to read message payload: %v", err)
	}

	var payload Payload
	json.Unmarshal(bytes, &payload)

	g.logger.Debugw("handling", "opCode", payload.Op)
	switch payload.Op {
	case OpCodeDispatch:
		g.handleOpCodeDispatch(payload)
	// todo handle the others
	case OpCodeHeartbeat:
		g.handleOpCodeHeartbeat(payload)
	case OpCodeIdentify:
		g.handleOpCodeIdentify(payload)
	case OpCodeReconnect:
		g.handleOpCodeReconnect(payload)
	case OpCodeInvalidSession:
		g.handleOpCodeInvalidSession(payload)
	case OpCodeHello:
		g.handleOpCodeHello(payload)
	case OpCodeHeartbeatACK:
		g.handleHeartBeatACK(payload)
	}

}
