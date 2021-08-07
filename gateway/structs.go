package gateway

import (
	"encoding/json"
)

type Payload struct {
	// Op [opcode](#DOCS_TOPICS_OPCODES_AND_STATUS_CODES/gateway-opcodes) for the payload
	Op int `json:"op"`
	// D event data
	D interface{} `json:"d"`
	// S sequence number, used for resuming sessions and heartbeats
	S int `json:"s"`
	// T the event name for this payload
	T string `json:"t"`
}

// MarshalData marshals the payload D to a provided struct value. v should always be a pointer to a target value
func (p *Payload) MarshalData(v interface{}) error {
	bytes, err := json.Marshal(p.D)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, v)
}

type Identify struct {
	// Token authentication token
	Token string `json:"token"`
	// Properties [connection properties](#DOCS_TOPICS_GATEWAY/identify-identify-connection-properties)
	Properties IdentifyConnectionProperties `json:"properties"`
	// Compress whether this connection supports compression of packets
	Compress bool `json:"compress"`
	// LargeThreshold value between 50 and 250, total number of members where the gateway will stop sending offline members in the guild member list
	LargeThreshold int `json:"large_threshold"`
	// Shard used for [Guild Sharding](#DOCS_TOPICS_GATEWAY/sharding)
	Shard []int `json:"shard"`
	// Presence presence structure for initial presence information
	Presence PresenceUpdateEventData `json:"presence"`
	// Intents the [Gateway Intents](#DOCS_TOPICS_GATEWAY/gateway-intents) you wish to receive
	Intents int `json:"intents"`
}

type IdentifyConnectionProperties struct {
	OS      string `json:"$os"`
	Browser string `json:"$browser"`
	Device  string `json:"$device"`
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

type Activity struct {
	// Name the activity's name
	Name string `json:"name"`
	// Type [activity type](#DOCS_TOPICS_GATEWAY/activity-object-activity-types)
	Type int `json:"type"`
	// Url stream url, is validated when type is 1
	Url string `json:"url"`
	// CreatedAt unix timestamp (in milliseconds) of when the activity was added to the user's session
	CreatedAt int `json:"created_at"`
	// Timestamps unix timestamps for start and/or end of the game
	Timestamps ActivityTimestamps `json:"timestamps"`
	// ApplicationId application id for the game
	ApplicationId string `json:"application_id"`
	// Details what the player is currently doing
	Details string `json:"details"`
	// State the user's current party status
	State string `json:"state"`
	// Emoji the emoji used for a custom status
	Emoji ActivityEmoji `json:"emoji"`
	// Party information for the current party of the player
	Party ActivityParty `json:"party"`
	// Assets images for the presence and their hover texts
	Assets ActivityAssets `json:"assets"`
	// Secrets secrets for Rich Presence joining and spectating
	Secrets ActivitySecrets `json:"secrets"`
	// Instance whether or not the activity is an instanced game session
	Instance bool `json:"instance"`
	// Flags [activity flags](#DOCS_TOPICS_GATEWAY/activity-object-activity-flags) `OR`d together, describes what the payload includes
	Flags int `json:"flags"`
	// Buttons the custom buttons shown in the Rich Presence (max 2)
	Buttons []ActivityButtons `json:"buttons"`
}

type ActivityButtons struct {
	// Label the text shown on the button (1-32 characters)
	Label string `json:"label"`
	// Url the url opened when clicking the button (1-512 characters)
	Url string `json:"url"`
}

type ActivityEmoji struct {
	Name     string `json:"name"`
	Id       string `json:"id"`
	Animated bool   `json:"animated"`
}

type ActivityParty struct {
	Id   string `json:"id"`
	Size []int  `json:"size"`
}

type ActivityTimestamps struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type ActivitySecrets struct {
	Join     string `json:"join"`
	Spectate string `json:"spectate"`
	Match    string `json:"match"`
}

type ActivityAssets struct {
	LargeImage string `json:"large_image"`
	LargeText  string `json:"large_text"`
	SmallImage string `json:"small_image"`
	SmallText  string `json:"small_text"`
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
