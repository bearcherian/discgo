package gateway

import (
	"time"
)

const (
	// EventHello defines the heartbeat interval
	EventHello = "HELLO"
	// EventReady contains the initial state information
	EventReady = "READY"
	// EventResumed response to Resume
	EventResumed = "RESUMED"
	// EventReconnect server is going away
	EventReconnect = "RECONNECT"
	// EventInvalidSession failure response to Identify or Resume or invalid active session
	EventInvalidSession = "INVALID_SESSION"
	// EventApplicationCommandCreate new Slash Command was created
	EventApplicationCommandCreate = "APPLICATION_COMMAND_CREATE"
	// EventApplicationCommandUpdate Slash Command was updated
	EventApplicationCommandUpdate = "APPLICATION_COMMAND_UPDATE"
	// EventApplicationCommandDelete Slash Command was deleted
	EventApplicationCommandDelete = "APPLICATION_COMMAND_DELETE"
	// EventChannelCreate new guild channel created
	EventChannelCreate = "CHANNEL_CREATE"
	// EventChannelUpdate channel was updated
	EventChannelUpdate = "CHANNEL_UPDATE"
	// EventChannelDelete channel was deleted
	EventChannelDelete = "CHANNEL_DELETE"
	// EventChannelPinsUpdate message was pinned or unpinned
	EventChannelPinsUpdate = "CHANNEL_PINS_UPDATE"
	// EventThreadCreate thread created
	EventThreadCreate = "THREAD_CREATE"
	// EventThreadUpdate thread was updated
	EventThreadUpdate = "THREAD_UPDATE"
	// EventThreadDelete thread was deleted
	EventThreadDelete = "THREAD_DELETE"
	// EventThreadListSync sent when gaining access to a channel
	EventThreadListSync = "THREAD_LIST_SYNC"
	// EventThreadMemberUpdate thread member for the current user was updated
	EventThreadMemberUpdate = "THREAD_MEMBER_UPDATE"
	// EventThreadMembersUpdate some user(s) were added to or removed from a thread
	EventThreadMembersUpdate = "THREAD_MEMBERS_UPDATE"
	// EventGuildCreate lazy-load for unavailable guild
	EventGuildCreate = "GUILD_CREATE"
	// EventGuildUpdate guild was updated
	EventGuildUpdate = "GUILD_UPDATE"
	// EventGuildDelete guild became unavailable
	EventGuildDelete = "GUILD_DELETE"
	// EventGuildBanAdd user was banned from a guild
	EventGuildBanAdd = "GUILD_BAN_ADD"
	// EventGuildBanRemove user was unbanned from a guild
	EventGuildBanRemove = "GUILD_BAN_REMOVE"
	// EventGuildEmojisUpdate guild emojis were updated
	EventGuildEmojisUpdate = "GUILD_EMOJIS_UPDATE"
	// EventGuildIntegrationsUpdate guild integration was updated
	EventGuildIntegrationsUpdate = "GUILD_INTEGRATIONS_UPDATE"
	// EventGuildMemberAdd new user joined a guild
	EventGuildMemberAdd = "GUILD_MEMBER_ADD"
	// EventGuildMemberRemove user was removed from a guild
	EventGuildMemberRemove = "GUILD_MEMBER_REMOVE"
	// EventGuildMemberUpdate guild member was updated
	EventGuildMemberUpdate = "GUILD_MEMBER_UPDATE"
	// EventGuildMembersChunk response to Request Guild Members
	EventGuildMembersChunk = "GUILD_MEMBERS_CHUNK"
	// EventGuildRoleCreate guild role was created
	EventGuildRoleCreate = "GUILD_ROLE_CREATE"
	// EventGuildRoleUpdate guild role was updated
	EventGuildRoleUpdate = "GUILD_ROLE_UPDATE"
	// EventGuildRoleDelete guild role was deleted
	EventGuildRoleDelete = "GUILD_ROLE_DELETE"
	// EventIntegrationCreate guild integration was created
	EventIntegrationCreate = "INTEGRATION_CREATE"
	// EventIntegrationUpdate guild integration was updated
	EventIntegrationUpdate = "INTEGRATION_UPDATE"
	// EventIntegrationDelete guild integration was deleted
	EventIntegrationDelete = "INTEGRATION_DELETE"
	// EventInteractionCreate user used an interaction
	EventInteractionCreate = "INTERACTION_CREATE"
	// EventInviteCreate invite to a channel was created
	EventInviteCreate = "INVITE_CREATE"
	// EventInviteDelete invite to a channel was deleted
	EventInviteDelete = "INVITE_DELETE"
	// EventMessageCreate message was created
	EventMessageCreate = "MESSAGE_CREATE"
	// EventMessageUpdate message was edited
	EventMessageUpdate = "MESSAGE_UPDATE"
	// EventMessageDelete message was deleted
	EventMessageDelete = "MESSAGE_DELETE"
	// EventMessageDeleteBulk multiple messages were deleted at once
	EventMessageDeleteBulk = "MESSAGE_DELETE_BULK"
	// EventMessageReactionAdd user reacted to a message
	EventMessageReactionAdd = "MESSAGE_REACTION_ADD"
	// EventMessageReactionRemove user removed a reaction from a message
	EventMessageReactionRemove = "MESSAGE_REACTION_REMOVE"
	// EventMessageReactionRemoveAll all reactions were explicitly removed from a message
	EventMessageReactionRemoveAll = "MESSAGE_REACTION_REMOVE_ALL"
	// EventMessageReactionRemoveEmoji all reactions for a given emoji were explicitly removed from a message
	EventMessageReactionRemoveEmoji = "MESSAGE_REACTION_REMOVE_EMOJI"
	// EventPresenceUpdate user was updated
	EventPresenceUpdate = "PRESENCE_UPDATE"
	// EventStageInstanceCreate stage instance was created
	EventStageInstanceCreate = "STAGE_INSTANCE_CREATE"
	// EventStageInstanceDelete stage instance was deleted or closed
	EventStageInstanceDelete = "STAGE_INSTANCE_DELETE"
	// EventStageInstanceUpdate stage instance was updated
	EventStageInstanceUpdate = "STAGE_INSTANCE_UPDATE"
	// EventTypingStart user started typing in a channel
	EventTypingStart = "TYPING_START"
	// EventUserUpdate properties about the user changed
	EventUserUpdate = "USER_UPDATE"
	// EventVoiceStateUpdate someone joined
	EventVoiceStateUpdate = "VOICE_STATE_UPDATE"
	// EventVoiceServerUpdate guild's voice server was updated
	EventVoiceServerUpdate = "VOICE_SERVER_UPDATE"
	// EventWebhooksUpdate guild channel webhook was created
	EventWebhooksUpdate = "WEBHOOKS_UPDATE"
)

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
	Guilds []Guild `json:"guilds"`
	// SessionId used for resuming connections
	SessionId string `json:"session_id"`
	// Shard the [shard information](#DOCS_TOPICS_GATEWAY/sharding) associated with this session, if sent when identifying
	Shard []int `json:"shard"`
	// Application contains `id` and `flags`
	Application Application `json:"application"`
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
	Emojis []Emoji `json:"emojis"`
}

type GuildStickersUpdateEventData struct {
	// GuildId id of the guild
	GuildId string `json:"guild_id"`
	// Stickers array of [stickers](#DOCS_RESOURCES_STICKER/sticker-object)
	Stickers []Sticker `json:"stickers"`
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
	NotFound []string `json:"not_found"`
	// Presences if passing true to `REQUEST_GUILD_MEMBERS`, presences of the returned members will be here
	Presences []PresenceUpdateEventData `json:"presences"`
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
	Member GuildMember `json:"member"`
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
	ClientStatus ClientStatus `json:"client_status"`
}

// ClientStatus Active sessions are indicated with an "online", "idle", or "dnd" string per platform.
// If a user is offline or invisible, the corresponding field is not present.
type ClientStatus struct {
	Desktop string `json:"desktop"`
	Mobile  string `json:"mobile"`
	Web     string `json:"web"`
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
	Member GuildMember `json:"member"`
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

func (g *Gateway) handleEventReady(p Payload) {

}

func (g *Gateway) handleEventHello(p Payload) error {
	var helloData Hello
	if err := p.MarshalData(&helloData); err != nil {
		return err
	}

	g.heartbeatInterval = helloData.HeartbeatInterval

	return nil
}
