package structs

type GatewayPayload struct {
	// Op [opcode](#DOCS_TOPICS_OPCODES_AND_STATUS_CODES/gateway-opcodes) for the payload
	Op int `json:"op"`
	// D event data
	D mixed (any JSON value) `json:"d"`
	// S sequence number, used for resuming sessions and heartbeats
	S int `json:"s"`
	// T the event name for this payload
	T string `json:"t"`
}

type Identify struct {
	// Token authentication token
	Token string `json:"token"`
	// Properties [connection properties](#DOCS_TOPICS_GATEWAY/identify-identify-connection-properties)
	Properties  `json:"properties"`
	// Compress whether this connection supports compression of packets
	Compress bool `json:"compress"`
	// LargeThreshold value between 50 and 250, total number of members where the gateway will stop sending offline members in the guild member list
	LargeThreshold int `json:"large_threshold"`
	// Shard used for [Guild Sharding](#DOCS_TOPICS_GATEWAY/sharding)
	Shard array of two integers (shard_id, num_shards) `json:"shard"`
	// Presence presence structure for initial presence information
	Presence UpdatePresence `json:"presence"`
	// Intents the [Gateway Intents](#DOCS_TOPICS_GATEWAY/gateway-intents) you wish to receive
	Intents int `json:"intents"`
}

type Resume struct {
	// Token session token
	Token string `json:"token"`
	// SessionId session id
	SessionId string `json:"session_id"`
	// Seq last sequence number received
	Seq int `json:"seq"`
}

type GuildRequestMembers struct {
	// GuildId id of the guild to get members for
	GuildId string `json:"guild_id"`
	// Query string that username starts with, or an empty string to return all members
	Query string `json:"query"`
	// Limit maximum number of members to send matching the `query`; a limit of `0` can be used with an empty string `query` to return all members
	Limit int `json:"limit"`
	// Presences used to specify if we want the presences of the matched members
	Presences bool `json:"presences"`
	// UserIds used to specify which users you wish to fetch
	UserIds string `json:"user_ids"`
	// Nonce nonce to identify the [Guild Members Chunk](#DOCS_TOPICS_GATEWAY/guild-members-chunk) response
	Nonce string `json:"nonce"`
}

type GatewayVoiceStateUpdate struct {
	// GuildId id of the guild
	GuildId string `json:"guild_id"`
	// ChannelId id of the voice channel client wants to join (null if disconnecting)
	ChannelId string `json:"channel_id"`
	// SelfMute is the client muted
	SelfMute bool `json:"self_mute"`
	// SelfDeaf is the client deafened
	SelfDeaf bool `json:"self_deaf"`
}

type GatewayPresenceUpdate struct {
	// Since unix time (in milliseconds) of when the client went idle, or null if the client is not idle
	Since int `json:"since"`
	// Activities the user's activities
	Activities []Activity `json:"activities"`
	// Status the user's new [status](#DOCS_TOPICS_GATEWAY/update-status-status-types)
	Status string `json:"status"`
	// Afk whether or not the client is afk
	Afk bool `json:"afk"`
}

type Hello struct {
	// HeartbeatInterval the interval (in milliseconds) the client should heartbeat with
	HeartbeatInterval int `json:"heartbeat_interval"`
}

type SessionStartLimit struct {
	// Total The total number of session starts the current user is allowed
	Total int `json:"total"`
	// Remaining The remaining number of session starts the current user is allowed
	Remaining int `json:"remaining"`
	// ResetAfter The number of milliseconds after which the limit resets
	ResetAfter int `json:"reset_after"`
	// MaxConcurrency The number of identify requests allowed per 5 seconds
	MaxConcurrency int `json:"max_concurrency"`
}