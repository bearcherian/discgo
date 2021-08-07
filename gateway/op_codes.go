package gateway

const (
	OpCodeDispatch            = 0
	OpCodeHeartbeat           = 1
	OpCodeIdentify            = 2
	OpCodePresenceUpdate      = 3
	OpCodeVoiceStateUpdate    = 4
	OpCodeResume              = 6
	OpCodeReconnect           = 7
	OpCodeRequestGuildMembers = 8
	OpCodeInvalidSession      = 9
	OpCodeHello               = 10
	OpCodeHeartbeatACK        = 11
)

func (g *Gateway) handleOpCodeDispatch(payload Payload) {

	// update the last sequence we recieved. only sent on Dispatch messages
	if payload.S > g.lastSequence {
		g.lastSequence = payload.S
	}

	switch payload.T {
	case EventReady:
		g.handleEventReady(payload)
	default:
		// all other event handlers are customizable event handlers
		handlers, found := g.eventHandlers[payload.T]
		if found {
			for _, handler := range handlers {
				handler(payload.D)
			}
		} else {
			g.logger.Debugw("no handlers found", "T", payload.T)
		}
	}
}

func (g *Gateway) handleOpCodeHeartbeat(p Payload) {
}
func (g *Gateway) handleOpCodeIdentify(p Payload) {
}
func (g *Gateway) handleOpCodeReconnect(p Payload) {

}
func (g *Gateway) handleOpCodeInvalidSession(p Payload) {

}
func (g *Gateway) handleOpCodeHello(p Payload) {

}
func (g *Gateway) handleHeartBeatACK(p Payload) {

}
