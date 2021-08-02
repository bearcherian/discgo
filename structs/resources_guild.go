package structs

type Guild struct {
	// Id guild id
	Id string `json:"id"`
	// Name guild name (2-100 characters, excluding trailing and leading whitespace)
	Name string `json:"name"`
	// Icon [icon hash](#DOCS_REFERENCE/image-formatting)
	Icon string `json:"icon"`
	// IconHash [icon hash](#DOCS_REFERENCE/image-formatting), returned when in the template object
	IconHash string `json:"icon_hash"`
	// Splash [splash hash](#DOCS_REFERENCE/image-formatting)
	Splash string `json:"splash"`
	// DiscoverySplash [discovery splash hash](#DOCS_REFERENCE/image-formatting); only present for guilds with the "DISCOVERABLE" feature
	DiscoverySplash string `json:"discovery_splash"`
	// Owner true if [the user](#DOCS_RESOURCES_USER/get-current-user-guilds) is the owner of the guild
	Owner bool `json:"owner"`
	// OwnerId id of owner
	OwnerId string `json:"owner_id"`
	// Permissions total permissions for [the user](#DOCS_RESOURCES_USER/get-current-user-guilds) in the guild (excludes overwrites)
	Permissions string `json:"permissions"`
	// Region [voice region](#DOCS_RESOURCES_VOICE/voice-region-object) id for the guild (deprecated)
	Region string `json:"region"`
	// AfkChannelId id of afk channel
	AfkChannelId string `json:"afk_channel_id"`
	// AfkTimeout afk timeout in seconds
	AfkTimeout int `json:"afk_timeout"`
	// WidgetEnabled true if the server widget is enabled
	WidgetEnabled bool `json:"widget_enabled"`
	// WidgetChannelId the channel id that the widget will generate an invite to, or `null` if set to no invite
	WidgetChannelId string `json:"widget_channel_id"`
	// VerificationLevel [verification level](#DOCS_RESOURCES_GUILD/guild-object-verification-level) required for the guild
	VerificationLevel int `json:"verification_level"`
	// DefaultMessageNotifications default [message notifications level](#DOCS_RESOURCES_GUILD/guild-object-default-message-notification-level)
	DefaultMessageNotifications int `json:"default_message_notifications"`
	// ExplicitContentFilter [explicit content filter level](#DOCS_RESOURCES_GUILD/guild-object-explicit-content-filter-level)
	ExplicitContentFilter int `json:"explicit_content_filter"`
	// Roles roles in the guild
	Roles []Role `json:"roles"`
	// Emojis custom guild emojis
	Emojis []Emoji `json:"emojis"`
	// Features enabled guild features
	Features []string `json:"features"`
	// MfaLevel required [MFA level](#DOCS_RESOURCES_GUILD/guild-object-mfa-level) for the guild
	MfaLevel int `json:"mfa_level"`
	// ApplicationId application id of the guild creator if it is bot-created
	ApplicationId string `json:"application_id"`
	// SystemChannelId the id of the channel where guild notices such as welcome messages and boost events are posted
	SystemChannelId string `json:"system_channel_id"`
	// SystemChannelFlags [system channel flags](#DOCS_RESOURCES_GUILD/guild-object-system-channel-flags)
	SystemChannelFlags int `json:"system_channel_flags"`
	// RulesChannelId the id of the channel where Community guilds can display rules and/or guidelines
	RulesChannelId string `json:"rules_channel_id"`
	// JoinedAt when this guild was joined at
	JoinedAt time.Time `json:"joined_at"`
	// Large true if this is considered a large guild
	Large bool `json:"large"`
	// Unavailable true if this guild is unavailable due to an outage
	Unavailable bool `json:"unavailable"`
	// MemberCount total number of members in this guild
	MemberCount int `json:"member_count"`
	// VoiceStates states of members currently in voice channels; lacks the `guild_id` key
	VoiceStates []VoiceState `json:"voice_states"`
	// Members users in the guild
	Members []GuildMember `json:"members"`
	// Channels channels in the guild
	Channels []Channel `json:"channels"`
	// Threads all active threads in the guild that current user has permission to view
	Threads []Channel `json:"threads"`
	// Presences presences of the members in the guild, will only include non-offline members if the size is greater than `large threshold`
	Presences []PresenceUpdate `json:"presences"`
	// MaxPresences the maximum number of presences for the guild (`null` is always returned, apart from the largest of guilds)
	MaxPresences int `json:"max_presences"`
	// MaxMembers the maximum number of members for the guild
	MaxMembers int `json:"max_members"`
	// VanityUrlCode the vanity url code for the guild
	VanityUrlCode string `json:"vanity_url_code"`
	// Description the description of a Community guild
	Description string `json:"description"`
	// Banner [banner hash](#DOCS_REFERENCE/image-formatting)
	Banner string `json:"banner"`
	// PremiumTier [premium tier](#DOCS_RESOURCES_GUILD/guild-object-premium-tier) (Server Boost level)
	PremiumTier int `json:"premium_tier"`
	// PremiumSubscriptionCount the number of boosts this guild currently has
	PremiumSubscriptionCount int `json:"premium_subscription_count"`
	// PreferredLocale the preferred locale of a Community guild; used in server discovery and notices from Discord; defaults to "en-US"
	PreferredLocale string `json:"preferred_locale"`
	// PublicUpdatesChannelId the id of the channel where admins and moderators of Community guilds receive notices from Discord
	PublicUpdatesChannelId string `json:"public_updates_channel_id"`
	// MaxVideoChannelUsers the maximum amount of users in a video channel
	MaxVideoChannelUsers int `json:"max_video_channel_users"`
	// ApproximateMemberCount approximate number of members in this guild, returned from the `GET /guilds/<id>` endpoint when `with_counts` is `true`
	ApproximateMemberCount int `json:"approximate_member_count"`
	// ApproximatePresenceCount approximate number of non-offline members in this guild, returned from the `GET /guilds/<id>` endpoint when `with_counts` is `true`
	ApproximatePresenceCount int `json:"approximate_presence_count"`
	// WelcomeScreen the welcome screen of a Community guild, shown to new members, returned in an [Invite](#DOCS_RESOURCES_INVITE/invite-object)'s guild object
	WelcomeScreen WelcomeScreen `json:"welcome_screen"`
	// NsfwLevel [guild NSFW level](#DOCS_RESOURCES_GUILD/guild-object-guild-nsfw-level)
	NsfwLevel int `json:"nsfw_level"`
	// StageInstances Stage instances in the guild
	StageInstances []StageInstance `json:"stage_instances"`
	// Stickers custom guild stickers
	Stickers []Sticker `json:"stickers"`
}

type GuildPreview struct {
	// Id guild id
	Id string `json:"id"`
	// Name guild name (2-100 characters)
	Name string `json:"name"`
	// Icon [icon hash](#DOCS_REFERENCE/image-formatting)
	Icon string `json:"icon"`
	// Splash [splash hash](#DOCS_REFERENCE/image-formatting)
	Splash string `json:"splash"`
	// DiscoverySplash [discovery splash hash](#DOCS_REFERENCE/image-formatting)
	DiscoverySplash string `json:"discovery_splash"`
	// Emojis custom guild emojis
	Emojis []Emoji `json:"emojis"`
	// Features enabled guild features
	Features []string `json:"features"`
	// ApproximateMemberCount approximate number of members in this guild
	ApproximateMemberCount int `json:"approximate_member_count"`
	// ApproximatePresenceCount approximate number of online members in this guild
	ApproximatePresenceCount int `json:"approximate_presence_count"`
	// Description the description for the guild, if the guild is discoverable
	Description string `json:"description"`
}

type GuildWidget struct {
	// Enabled whether the widget is enabled
	Enabled bool `json:"enabled"`
	// ChannelId the widget channel id
	ChannelId string `json:"channel_id"`
}

type GuildMember struct {
	// User the user this guild member represents
	User User `json:"user"`
	// Nick this users guild nickname
	Nick string `json:"nick"`
	// Roles array of [role](#DOCS_TOPICS_PERMISSIONS/role-object) object ids
	Roles []string `json:"roles"`
	// JoinedAt when the user joined the guild
	JoinedAt time.Time `json:"joined_at"`
	// PremiumSince when the user started [boosting](https://support.discord.com/hc/en-us/articles/360028038352-Server-Boosting-) the guild
	PremiumSince time.Time `json:"premium_since"`
	// Deaf whether the user is deafened in voice channels
	Deaf bool `json:"deaf"`
	// Mute whether the user is muted in voice channels
	Mute bool `json:"mute"`
	// Pending whether the user has not yet passed the guild's [Membership Screening](#DOCS_RESOURCES_GUILD/membership-screening-object) requirements
	Pending bool `json:"pending"`
	// Permissions total permissions of the member in the channel, including overwrites, returned when in the interaction object
	Permissions string `json:"permissions"`
}

type Integration struct {
	// Id integration id
	Id string `json:"id"`
	// Name integration name
	Name string `json:"name"`
	// Type integration type (twitch, youtube, or discord)
	Type string `json:"type"`
	// Enabled is this integration enabled
	Enabled bool `json:"enabled"`
	// Syncing is this integration syncing
	Syncing bool `json:"syncing"`
	// RoleId id that this integration uses for "subscribers"
	RoleId string `json:"role_id"`
	// EnableEmoticons whether emoticons should be synced for this integration (twitch only currently)
	EnableEmoticons bool `json:"enable_emoticons"`
	// ExpireBehavior the behavior of expiring subscribers
	ExpireBehavior [integration expire behavior](#DOCS_RESOURCES_GUILD/integration-object-integration-expire-behaviors) `json:"expire_behavior"`
	// ExpireGracePeriod the grace period (in days) before expiring subscribers
	ExpireGracePeriod int `json:"expire_grace_period"`
	// User user for this integration
	User User `json:"user"`
	// Account integration account information
	Account Account `json:"account"`
	// SyncedAt when this integration was last synced
	SyncedAt time.Time `json:"synced_at"`
	// SubscriberCount how many subscribers this integration has
	SubscriberCount int `json:"subscriber_count"`
	// Revoked has this integration been revoked
	Revoked bool `json:"revoked"`
	// Application The bot/OAuth2 application for discord integrations
	Application Application `json:"application"`
}

type IntegrationAccount struct {
	// Id id of the account
	Id string `json:"id"`
	// Name name of the account
	Name string `json:"name"`
}

type IntegrationApplication struct {
	// Id the id of the app
	Id string `json:"id"`
	// Name the name of the app
	Name string `json:"name"`
	// Icon the [icon hash](#DOCS_REFERENCE/image-formatting) of the app
	Icon string `json:"icon"`
	// Description the description of the app
	Description string `json:"description"`
	// Summary the summary of the app
	Summary string `json:"summary"`
	// Bot the bot associated with this application
	Bot User `json:"bot"`
}

type Ban struct {
	// Reason the reason for the ban
	Reason string `json:"reason"`
	// User the banned user
	User User `json:"user"`
}

type WelcomeScreen struct {
	// Description the server description shown in the welcome screen
	Description string `json:"description"`
	// WelcomeChannels the channels shown in the welcome screen, up to 5
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels"`
}

type WelcomeScreenChannel struct {
	// ChannelId the channel's id
	ChannelId string `json:"channel_id"`
	// Description the description shown for the channel
	Description string `json:"description"`
	// EmojiId the [emoji id](#DOCS_REFERENCE/image-formatting), if the emoji is custom
	EmojiId string `json:"emoji_id"`
	// EmojiName the emoji name if custom, the unicode character if standard, or `null` if no emoji is set
	EmojiName string `json:"emoji_name"`
}