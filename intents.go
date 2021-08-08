package discgo

const (
	// Guild
	IntentGuildCreate = iota + 1<<0
	IntentGuildUpdate
	IntentGuildDelete
	IntentGuildRoleCreate
	IntentGuildRoleUpdate
	IntentGuildRoleDelete
	IntentChannelCreate
	IntentChannelUpdate
	IntentChannelDelete
	IntentChannelPinsUpdate
	IntentThreadCreate
	IntentThreadUpdate
	IntentThreadDelete
	IntentThreadListSync
	IntentThreadMemberUpdate
	IntentGuildThreadMembersUpdate
	IntentStageInstanceCreate
	IntentStageInstanceUpdate
	IntentStageInstanceDelete

	// Guild Members
	IntentGuildMemberAdd = iota + 1<<1
	IntentGuildMemberUpdate
	IntentGuildMemberRemove
	IntentMembersThreadMembersUpdate

	// Guild Bans
	IntentGuildBanAdd = iota + 1<<2
	IntentGuildBanRemove

	//Guild Emojis
	IntentGuildEmojisUpdate = iota + 1<<3

	// Guild Integrations
	IntentGuildintegrationsUpdate = iota + 1<<4
	IntentIntegrationCreate
	IntentIntegrationUpdate
	IntentIntegrationDelete

	// Guild Webhooks
	IntentWebhooksUpdate = iota + 1<<5

	// Guild Invites
	IntentInviteCreate = iota + 1<<6
	IntentInviteDelete

	// Guild Voice States
	IntentVoiceStateUpdate = iota + 1<<7

	// Guild Presences
	IntentPresenceUpdate = iota + 1<<8

	// Guild Messages
	IntentMessageCreate = iota + 1<<9
	IntentMessageUpdate
	IntentMessageDelete
	IntentMessageDeleteBulk

	// Guild Message Reactions
	IntentMessageReactionAdd = iota + 1<<10
	IntentMessageReactionRemove
	IntentMessageReactionRemoveAll
	IntentMessageReactionRemoveEmoji

	// Guild Message Typing
	IntentTypingStart = iota + 1<<11

	// Direct Messages
	IntentDirectMessageCreate = iota + 1<<12
	IntentDirectMessageUpdate
	IntentDirectMessageDelete
	IntentDirectMessageChannelPinsUpdate

	// Direct Message Reactions
	IntentDirectMessageReactionAdd = iota + 1<<13
	IntentDirectMessageReactionRemove
	IntentDirectMessageReactionRemoveAll
	IntentDirectMessageReactionRemoveEmoji

	// Direct Message Typing
	IntentDirectMessageTypingStart = iota + 1<<14
)
