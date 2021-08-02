package structs

type VoiceState struct {
	// GuildId the guild id this voice state is for
	GuildId string `json:"guild_id"`
	// ChannelId the channel id this user is connected to
	ChannelId string `json:"channel_id"`
	// UserId the user id this voice state is for
	UserId string `json:"user_id"`
	// Member the guild member this voice state is for
	Member GuildMember `json:"member"`
	// SessionId the session id for this voice state
	SessionId string `json:"session_id"`
	// Deaf whether this user is deafened by the server
	Deaf bool `json:"deaf"`
	// Mute whether this user is muted by the server
	Mute bool `json:"mute"`
	// SelfDeaf whether this user is locally deafened
	SelfDeaf bool `json:"self_deaf"`
	// SelfMute whether this user is locally muted
	SelfMute bool `json:"self_mute"`
	// SelfStream whether this user is streaming using "Go Live"
	SelfStream bool `json:"self_stream"`
	// SelfVideo whether this user's camera is enabled
	SelfVideo bool `json:"self_video"`
	// Suppress whether this user is muted by the current user
	Suppress bool `json:"suppress"`
	// RequestToSpeakTimestamp the time at which the user requested to speak
	RequestToSpeakTimestamp time.Time `json:"request_to_speak_timestamp"`
}

type VoiceRegion struct {
	// Id unique ID for the region
	Id string `json:"id"`
	// Name name of the region
	Name string `json:"name"`
	// Vip true if this is a vip-only server
	Vip bool `json:"vip"`
	// Optimal true for a single server that is closest to the current user's client
	Optimal bool `json:"optimal"`
	// Deprecated whether this is a deprecated voice region (avoid switching to these)
	Deprecated bool `json:"deprecated"`
	// Custom whether this is a custom voice region (used for events/etc)
	Custom bool `json:"custom"`
}