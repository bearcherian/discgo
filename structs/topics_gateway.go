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

type ReadyEventData struct {
	// V [gateway version](#DOCS_TOPICS_GATEWAY/gateways-gateway-versions)
	V int `json:"v"`
	// User information about the user including email
	User User `json:"user"`
	// Guilds the guilds the user is in
	Guilds []UnavailableGuild `json:"guilds"`
	// SessionId used for resuming connections
	SessionId string `json:"session_id"`
	// Shard the [shard information](#DOCS_TOPICS_GATEWAY/sharding) associated with this session, if sent when identifying
	Shard array of two integers (shard_id, num_shards) `json:"shard"`
	// Application contains `id` and `flags`
	Application partial [application object](#DOCS_RESOURCES_APPLICATION/application-object) `json:"application"`
}

type ThreadListSyncEventData struct {
	// GuildId the id of the guild
	GuildId string `json:"guild_id"`
	// ChannelIds the parent channel ids whose threads are being synced.  If omitted, then threads were synced for the entire guild.  This array may contain channel_ids that have no active threads as well, so you know to clear that data.
	ChannelIds []string `json:"channel_ids"`
	// Threads all active threads in the given channels that the current user can access
	Threads []Channel `json:"threads"`
	// Members all thread member objects from the synced threads for the current user, indicating which threads the current user has been added to
	Members []ThreadMember `json:"members"`
}

type ThreadMembersUpdateEventData struct {
	// Id the id of the thread
	Id string `json:"id"`
	// GuildId the id of the guild
	GuildId string `json:"guild_id"`
	// MemberCount the approximate number of members in the thread, capped at 50
	MemberCount int `json:"member_count"`
	// AddedMembers the users who were added to the thread
	AddedMembers []ThreadMember `json:"added_members"`
	// RemovedMemberIds the id of the users who were removed from the thread
	RemovedMemberIds []string `json:"removed_member_ids"`
}

type ChannelPinsUpdateEventData struct {
	// GuildId the id of the guild
	GuildId string `json:"guild_id"`
	// ChannelId the id of the channel
	ChannelId string `json:"channel_id"`
	// LastPinTimestamp the time at which the most recent pinned message was pinned
	LastPinTimestamp time.Time `json:"last_pin_timestamp"`
}

type GuildBanAddEventData struct {
	// GuildId id of the guild
	GuildId string `json:"guild_id"`
	// User the banned user
	User User `json:"user"`
}

type GuildBanRemoveEventData struct {
	// GuildId id of the guild
	GuildId string `json:"guild_id"`
	// User the unbanned user
	User User `json:"user"`
}

type GuildEmojisUpdateEventData struct {
	// GuildId id of the guild
	GuildId string `json:"guild_id"`
	// Emojis array of [emojis](#DOCS_RESOURCES_EMOJI/emoji-object)
	Emojis array `json:"emojis"`
}

type GuildStickersUpdateEventData struct {
	// GuildId id of the guild
	GuildId string `json:"guild_id"`
	// Stickers array of [stickers](#DOCS_RESOURCES_STICKER/sticker-object)
	Stickers array `json:"stickers"`
}

type GuildIntegrationsUpdateEventData struct {
	// GuildId id of the guild whose integrations were updated
	GuildId string `json:"guild_id"`
}

type GuildMemberRemoveEventData struct {
	// GuildId the id of the guild
	GuildId string `json:"guild_id"`
	// User the user who was removed
	User User `json:"user"`
}

type GuildMemberUpdateEventData struct {
	// GuildId the id of the guild
	GuildId string `json:"guild_id"`
	// Roles user role ids
	Roles []string `json:"roles"`
	// User the user
	User User `json:"user"`
	// Nick nickname of the user in the guild
	Nick string `json:"nick"`
	// JoinedAt when the user joined the guild
	JoinedAt time.Time `json:"joined_at"`
	// PremiumSince when the user starting [boosting](https://support.discord.com/hc/en-us/articles/360028038352-Server-Boosting-) the guild
	PremiumSince time.Time `json:"premium_since"`
	// Deaf whether the user is deafened in voice channels
	Deaf bool `json:"deaf"`
	// Mute whether the user is muted in voice channels
	Mute bool `json:"mute"`
	// Pending whether the user has not yet passed the guild's [Membership Screening](#DOCS_RESOURCES_GUILD/membership-screening-object) requirements
	Pending bool `json:"pending"`
}

type GuildMembersChunkEventData struct {
	// GuildId the id of the guild
	GuildId string `json:"guild_id"`
	// Members set of guild members
	Members []GuildMember `json:"members"`
	// ChunkIndex the chunk index in the expected chunks for this response (0 <= chunk\_index < chunk\_count)
	ChunkIndex int `json:"chunk_index"`
	// ChunkCount the total number of expected chunks for this response
	ChunkCount int `json:"chunk_count"`
	// NotFound if passing an invalid id to `REQUEST_GUILD_MEMBERS`, it will be returned here
	NotFound array `json:"not_found"`
	// Presences if passing true to `REQUEST_GUILD_MEMBERS`, presences of the returned members will be here
	Presences []Presence `json:"presences"`
	// Nonce the nonce used in the [Guild Members Request](#DOCS_TOPICS_GATEWAY/request-guild-members)
	Nonce string `json:"nonce"`
}

type GuildRoleCreateEventData struct {
	// GuildId the id of the guild
	GuildId string `json:"guild_id"`
	// Role the role created
	Role Role `json:"role"`
}

type GuildRoleUpdateEventData struct {
	// GuildId the id of the guild
	GuildId string `json:"guild_id"`
	// Role the role updated
	Role Role `json:"role"`
}

type GuildRoleDeleteEventData struct {
	// GuildId id of the guild
	GuildId string `json:"guild_id"`
	// RoleId id of the role
	RoleId string `json:"role_id"`
}

type IntegrationDeleteEventData struct {
	// Id integration id
	Id string `json:"id"`
	// GuildId id of the guild
	GuildId string `json:"guild_id"`
	// ApplicationId id of the bot/OAuth2 application for this discord integration
	ApplicationId string `json:"application_id"`
}

type InviteCreateEventData struct {
	// ChannelId the channel the invite is for
	ChannelId string `json:"channel_id"`
	// Code the unique invite [code](#DOCS_RESOURCES_INVITE/invite-object)
	Code string `json:"code"`
	// CreatedAt the time at which the invite was created
	CreatedAt time.Time `json:"created_at"`
	// GuildId the guild of the invite
	GuildId string `json:"guild_id"`
	// Inviter the user that created the invite
	Inviter User `json:"inviter"`
	// MaxAge how long the invite is valid for (in seconds)
	MaxAge int `json:"max_age"`
	// MaxUses the maximum number of times the invite can be used
	MaxUses int `json:"max_uses"`
	// TargetType the [type of target](#DOCS_RESOURCES_INVITE/invite-object-invite-target-types) for this voice channel invite
	TargetType int `json:"target_type"`
	// TargetUser the user whose stream to display for this voice channel stream invite
	TargetUser User `json:"target_user"`
	// TargetApplication the embedded application to open for this voice channel embedded application invite
	TargetApplication Application `json:"target_application"`
	// Temporary whether or not the invite is temporary (invited users will be kicked on disconnect unless they're assigned a role)
	Temporary bool `json:"temporary"`
	// Uses how many times the invite has been used (always will be 0)
	Uses int `json:"uses"`
}

type InviteDeleteEventData struct {
	// ChannelId the channel of the invite
	ChannelId string `json:"channel_id"`
	// GuildId the guild of the invite
	GuildId string `json:"guild_id"`
	// Code the unique invite [code](#DOCS_RESOURCES_INVITE/invite-object)
	Code string `json:"code"`
}

type MessageDeleteEventData struct {
	// Id the id of the message
	Id string `json:"id"`
	// ChannelId the id of the channel
	ChannelId string `json:"channel_id"`
	// GuildId the id of the guild
	GuildId string `json:"guild_id"`
}

type MessageDeleteBulkEventData struct {
	// Ids the ids of the messages
	Ids []string `json:"ids"`
	// ChannelId the id of the channel
	ChannelId string `json:"channel_id"`
	// GuildId the id of the guild
	GuildId string `json:"guild_id"`
}

type MessageReactionAddEventData struct {
	// UserId the id of the user
	UserId string `json:"user_id"`
	// ChannelId the id of the channel
	ChannelId string `json:"channel_id"`
	// MessageId the id of the message
	MessageId string `json:"message_id"`
	// GuildId the id of the guild
	GuildId string `json:"guild_id"`
	// Member the member who reacted if this happened in a guild
	Member Member `json:"member"`
	// Emoji the emoji used to react - [example](#DOCS_RESOURCES_EMOJI/emoji-object-gateway-reaction-standard-emoji-example)
	Emoji Emoji `json:"emoji"`
}

type MessageReactionRemoveEventData struct {
	// UserId the id of the user
	UserId string `json:"user_id"`
	// ChannelId the id of the channel
	ChannelId string `json:"channel_id"`
	// MessageId the id of the message
	MessageId string `json:"message_id"`
	// GuildId the id of the guild
	GuildId string `json:"guild_id"`
	// Emoji the emoji used to react - [example](#DOCS_RESOURCES_EMOJI/emoji-object-gateway-reaction-standard-emoji-example)
	Emoji Emoji `json:"emoji"`
}

type MessageReactionRemoveAllEventData struct {
	// ChannelId the id of the channel
	ChannelId string `json:"channel_id"`
	// MessageId the id of the message
	MessageId string `json:"message_id"`
	// GuildId the id of the guild
	GuildId string `json:"guild_id"`
}

type PresenceUpdateEventData struct {
	// User the user presence is being updated for
	User User `json:"user"`
	// GuildId id of the guild
	GuildId string `json:"guild_id"`
	// Status either "idle", "dnd", "online", or "offline"
	Status string `json:"status"`
	// Activities user's current activities
	Activities []Activity `json:"activities"`
	// ClientStatus user's platform-dependent status
	ClientStatus Client_status `json:"client_status"`
}

type TypingStartEventData struct {
	// ChannelId id of the channel
	ChannelId string `json:"channel_id"`
	// GuildId id of the guild
	GuildId string `json:"guild_id"`
	// UserId id of the user
	UserId string `json:"user_id"`
	// Timestamp unix time (in seconds) of when the user started typing
	Timestamp int `json:"timestamp"`
	// Member the member who started typing if this happened in a guild
	Member Member `json:"member"`
}

type VoiceServerUpdateEventData struct {
	// Token voice connection token
	Token string `json:"token"`
	// GuildId the guild this voice server update is for
	GuildId string `json:"guild_id"`
	// Endpoint the voice server host
	Endpoint string `json:"endpoint"`
}

type WebhookUpdateEventData struct {
	// GuildId id of the guild
	GuildId string `json:"guild_id"`
	// ChannelId id of the channel
	ChannelId string `json:"channel_id"`
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