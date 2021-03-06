package discgo

import "time"

type AllowedMentionType string

const (
	AllowedMentionRole     AllowedMentionType = "roles"
	AllowedMentionUser     AllowedMentionType = "users"
	AllowedMentionEveryone AllowedMentionType = "everyone"
)

type AllowedMentions struct {
	// Parse An array of [allowed mention types](#DOCS_RESOURCES_CHANNEL/allowed-mentions-object-allowed-mention-types) to parse from the content.
	Parse []AllowedMentionType `json:"parse"`
	// Roles Array of role_ids to mention (Max size of 100)
	Roles string `json:"roles"`
	// Users Array of user_ids to mention (Max size of 100)
	Users string `json:"users"`
	// RepliedUser For replies, whether to mention the author of the message being replied to (default false)
	RepliedUser bool `json:"replied_user"`
}

type Application struct {
	// Id the id of the app
	Id string `json:"id"`
	// Name the name of the app
	Name string `json:"name"`
	// Icon the [icon hash](#DOCS_REFERENCE/image-formatting) of the app
	Icon string `json:"icon"`
	// Description the description of the app
	Description string `json:"description"`
	// RpcOrigins an array of rpc origin urls, if rpc is enabled
	RpcOrigins []string `json:"rpc_origins"`
	// BotPublic when false only app owner can join the app's bot to guilds
	BotPublic bool `json:"bot_public"`
	// BotRequireCodeGrant when true the app's bot will only join upon completion of the full oauth2 code grant flow
	BotRequireCodeGrant bool `json:"bot_require_code_grant"`
	// TermsOfServiceUrl the url of the app's terms of service
	TermsOfServiceUrl string `json:"terms_of_service_url"`
	// PrivacyPolicyUrl the url of the app's privacy policy
	PrivacyPolicyUrl string `json:"privacy_policy_url"`
	// Owner partial user object containing info on the owner of the application
	Owner User `json:"owner"`
	// Summary if this application is a game sold on Discord, this field will be the summary field for the store page of its primary sku
	Summary string `json:"summary"`
	// VerifyKey the hex encoded key for verification in interactions and the GameSDK's [GetTicket](#DOCS_GAME_SDK_APPLICATIONS/getticket)
	VerifyKey string `json:"verify_key"`
	// Team if the application belongs to a team, this will be a list of the members of that team
	Team Team `json:"team"`
	// GuildId if this application is a game sold on Discord, this field will be the guild to which it has been linked
	GuildId string `json:"guild_id"`
	// PrimarySkuId if this application is a game sold on Discord, this field will be the id of the "Game SKU" that is created, if exists
	PrimarySkuId string `json:"primary_sku_id"`
	// Slug if this application is a game sold on Discord, this field will be the URL slug that links to the store page
	Slug string `json:"slug"`
	// CoverImage the application's default rich presence invite [cover image hash](#DOCS_REFERENCE/image-formatting)
	CoverImage string `json:"cover_image"`
	// Flags the application's public [flags](#DOCS_RESOURCES_APPLICATION/application-object-application-flags)
	Flags int `json:"flags"`
}

type Attachment struct {
	// Id attachment id
	Id string `json:"id"`
	// Filename name of file attached
	Filename string `json:"filename"`
	// ContentType the attachment's [media type](https://en.wikipedia.org/wiki/Media_type)
	ContentType string `json:"content_type"`
	// Size size of file in bytes
	Size int `json:"size"`
	// Url source url of file
	Url string `json:"url"`
	// ProxyUrl a proxied url of file
	ProxyUrl string `json:"proxy_url"`
	// Height height of file (if image)
	Height int `json:"height"`
	// Width width of file (if image)
	Width int `json:"width"`
}

type Channel struct {
	// Id the id of this channel
	Id string `json:"id"`
	// Type the [type of channel](#DOCS_RESOURCES_CHANNEL/channel-object-channel-types)
	Type int `json:"type"`
	// GuildId the id of the guild (may be missing for some channel objects received over gateway guild dispatches)
	GuildId string `json:"guild_id"`
	// Position sorting position of the channel
	Position int `json:"position"`
	// PermissionOverwrites explicit permission overwrites for members and roles
	PermissionOverwrites []Overwrite `json:"permission_overwrites"`
	// Name the name of the channel (1-100 characters)
	Name string `json:"name"`
	// Topic the channel topic (0-1024 characters)
	Topic string `json:"topic"`
	// Nsfw whether the channel is nsfw
	Nsfw bool `json:"nsfw"`
	// LastMessageId the id of the last message sent in this channel (may not point to an existing or valid message)
	LastMessageId string `json:"last_message_id"`
	// Bitrate the bitrate (in bits) of the voice channel
	Bitrate int `json:"bitrate"`
	// UserLimit the user limit of the voice channel
	UserLimit int `json:"user_limit"`
	// RateLimitPerUser amount of seconds a user has to wait before sending another message (0-21600); bots, as well as users with the permission `manage_messages` or `manage_channel`, are unaffected
	RateLimitPerUser int `json:"rate_limit_per_user"`
	// Recipients the recipients of the DM
	Recipients []User `json:"recipients"`
	// Icon icon hash
	Icon string `json:"icon"`
	// OwnerId id of the creator of the group DM or thread
	OwnerId string `json:"owner_id"`
	// ApplicationId application id of the group DM creator if it is bot-created
	ApplicationId string `json:"application_id"`
	// ParentId for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created
	ParentId string `json:"parent_id"`
	// LastPinTimestamp when the last pinned message was pinned. This may be `null` in events such as `GUILD_CREATE` when a message is not pinned.
	LastPinTimestamp time.Time `json:"last_pin_timestamp"`
	// RtcRegion [voice region](#DOCS_RESOURCES_VOICE/voice-region-object) id for the voice channel, automatic when set to null
	RtcRegion string `json:"rtc_region"`
	// VideoQualityMode the camera [video quality mode](#DOCS_RESOURCES_CHANNEL/channel-object-video-quality-modes) of the voice channel, 1 when not present
	VideoQualityMode int `json:"video_quality_mode"`
	// MessageCount an approximate count of messages in a thread, stops counting at 50
	MessageCount int `json:"message_count"`
	// MemberCount an approximate count of users in a thread, stops counting at 50
	MemberCount int `json:"member_count"`
	// ThreadMetadata thread-specific fields not needed by other channels
	ThreadMetadata ThreadMetadata `json:"thread_metadata"`
	// Member thread member object for the current user, if they have joined the thread, only included on certain API endpoints
	Member ThreadMember `json:"member"`
	// DefaultAutoArchiveDuration default duration for newly created threads, in minutes, to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	DefaultAutoArchiveDuration int `json:"default_auto_archive_duration"`
	// Permissions computed permissions for the invoking user in the channel, including overwrites, only included when part of the `resolved` data received on a slash command interaction
	Permissions string `json:"permissions"`
}

type ChannelMention struct {
	// Id id of the channel
	Id string `json:"id"`
	// GuildId id of the guild containing the channel
	GuildId string `json:"guild_id"`
	// Type the [type of channel](#DOCS_RESOURCES_CHANNEL/channel-object-channel-types)
	Type int `json:"type"`
	// Name the name of the channel
	Name string `json:"name"`
}

type Component struct {
	// Type [component type](#DOCS_INTERACTIONS_MESSAGE_COMPONENTS/component-object-component-types)
	Type int `json:"type"`
	// CustomId a developer-defined identifier for the component, max 100 characters
	CustomId string `json:"custom_id"`
	// Disabled whether the component is disabled, default `false`
	Disabled bool `json:"disabled"`
	// Style one of [button styles](#DOCS_INTERACTIONS_MESSAGE_COMPONENTS/button-object-button-styles)
	Style int `json:"style"`
	// Label text that appears on the button, max 80 characters
	Label string `json:"label"`
	// Emoji `name`, `id`, and `animated`
	Emoji Emoji `json:"emoji"`
	// Url a url for link-style buttons
	Url string `json:"url"`
	// Options the choices in the select, max 25
	Options []SelectOption `json:"options"`
	// Placeholder custom placeholder text if nothing is selected, max 100 characters
	Placeholder string `json:"placeholder"`
	// MinValues the minimum number of items that must be chosen; default 1, min 0, max 25
	MinValues int `json:"min_values"`
	// MaxValues the maximum number of items that can be chosen; default 1, max 25
	MaxValues int `json:"max_values"`
	// Components a list of child components
	Components []Component `json:"components"`
}

type Embed struct {
	// Title title of embed
	Title string `json:"title"`
	// Type [type of embed](#DOCS_RESOURCES_CHANNEL/embed-object-embed-types) (always "rich" for webhook embeds)
	Type string `json:"type"`
	// Description description of embed
	Description string `json:"description"`
	// Url url of embed
	Url string `json:"url"`
	// Timestamp timestamp of embed content
	Timestamp time.Time `json:"timestamp"`
	// Color color code of the embed
	Color int `json:"color"`
	// Footer footer information
	Footer EmbedFooter `json:"footer"`
	// Image image information
	Image EmbedImage `json:"image"`
	// Thumbnail thumbnail information
	Thumbnail EmbedThumbnail `json:"thumbnail"`
	// Video video information
	Video EmbedVideo `json:"video"`
	// Provider provider information
	Provider EmbedProvider `json:"provider"`
	// Author author information
	Author EmbedAuthor `json:"author"`
	// Fields fields information
	Fields []EmbedField `json:"fields"`
}

type EmbedThumbnail struct {
	// Url source url of thumbnail (only supports http(s) and attachments)
	Url string `json:"url"`
	// ProxyUrl a proxied url of the thumbnail
	ProxyUrl string `json:"proxy_url"`
	// Height height of thumbnail
	Height int `json:"height"`
	// Width width of thumbnail
	Width int `json:"width"`
}

type EmbedVideo struct {
	// Url source url of video
	Url string `json:"url"`
	// ProxyUrl a proxied url of the video
	ProxyUrl string `json:"proxy_url"`
	// Height height of video
	Height int `json:"height"`
	// Width width of video
	Width int `json:"width"`
}

type EmbedImage struct {
	// Url source url of image (only supports http(s) and attachments)
	Url string `json:"url"`
	// ProxyUrl a proxied url of the image
	ProxyUrl string `json:"proxy_url"`
	// Height height of image
	Height int `json:"height"`
	// Width width of image
	Width int `json:"width"`
}

type EmbedProvider struct {
	// Name name of provider
	Name string `json:"name"`
	// Url url of provider
	Url string `json:"url"`
}

type EmbedAuthor struct {
	// Name name of author
	Name string `json:"name"`
	// Url url of author
	Url string `json:"url"`
	// IconUrl url of author icon (only supports http(s) and attachments)
	IconUrl string `json:"icon_url"`
	// ProxyIconUrl a proxied url of author icon
	ProxyIconUrl string `json:"proxy_icon_url"`
}

type EmbedFooter struct {
	// Text footer text
	Text string `json:"text"`
	// IconUrl url of footer icon (only supports http(s) and attachments)
	IconUrl string `json:"icon_url"`
	// ProxyIconUrl a proxied url of footer icon
	ProxyIconUrl string `json:"proxy_icon_url"`
}

type EmbedField struct {
	// Name name of the field
	Name string `json:"name"`
	// Value value of the field
	Value string `json:"value"`
	// Inline whether or not this field should display inline
	Inline bool `json:"inline"`
}

type Emoji struct {
	// Id [emoji id](#DOCS_REFERENCE/image-formatting)
	Id string `json:"id"`
	// Name emoji name
	Name string `json:"name"`
	// Roles roles allowed to use this emoji
	Roles []string `json:"roles"`
	// User user that created this emoji
	User User `json:"user"`
	// RequireColons whether this emoji must be wrapped in colons
	RequireColons bool `json:"require_colons"`
	// Managed whether this emoji is managed
	Managed bool `json:"managed"`
	// Animated whether this emoji is animated
	Animated bool `json:"animated"`
	// Available whether this emoji can be used, may be false due to loss of Server Boosts
	Available bool `json:"available"`
}

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
	Presences []PresenceUpdateEventData `json:"presences"`
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

type InteractionCallbackType int

const (
	InteractionCallbackPong                             InteractionCallbackType = 1
	InteractionCallbackChannelMessageWithSource         InteractionCallbackType = 4
	InteractionCallbackDeferredChannelMessageWithSource InteractionCallbackType = 5
	InteractionCallbackDeferredUpdateMessage            InteractionCallbackType = 6
	InteractionCallbackUpdateMessage                    InteractionCallbackType = 7
)

type InteractionResponse struct {
	// Type the type of response
	Type InteractionCallbackType `json:"type"`
	// Data an optional response message
	Data InteractionApplicationCommandCallbackData `json:"data"`
}

type InteractionApplicationCommandCallbackData struct {
	// Tts is the response TTS
	Tts bool `json:"tts"`
	// Content message content
	Content string `json:"content"`
	// Embeds supports up to 10 embeds
	Embeds []Embed `json:"embeds"`
	// AllowedMentions [allowed mentions](#DOCS_RESOURCES_CHANNEL/allowed-mentions-object) object
	AllowedMentions AllowedMentions `json:"allowed_mentions"`
	// Flags [interaction application command callback data flags](#DOCS_INTERACTIONS_SLASH_COMMANDS/interaction-response-object-interaction-application-command-callback-data-flags)
	Flags int `json:"flags"`
	// Components message components
	Components []Component `json:"components"`
}

type Message struct {
	// Id id of the message
	Id string `json:"id"`
	// ChannelId id of the channel the message was sent in
	ChannelId string `json:"channel_id"`
	// GuildId id of the guild the message was sent in
	GuildId string `json:"guild_id"`
	// Author the author of this message (not guaranteed to be a valid user, see below)
	Author User `json:"author"`
	// Member member properties for this message's author
	Member GuildMember `json:"member"`
	// Content contents of the message
	Content string `json:"content"`
	// Timestamp when this message was sent
	Timestamp time.Time `json:"timestamp"`
	// EditedTimestamp when this message was edited (or null if never)
	EditedTimestamp time.Time `json:"edited_timestamp"`
	// Tts whether this was a TTS message
	Tts bool `json:"tts"`
	// MentionEveryone whether this message mentions everyone
	MentionEveryone bool `json:"mention_everyone"`
	// Mentions users specifically mentioned in the message
	Mentions []User `json:"mentions"`
	// MentionRoles roles specifically mentioned in this message
	MentionRoles []string `json:"mention_roles"`
	// MentionChannels channels specifically mentioned in this message
	MentionChannels []ChannelMention `json:"mention_channels"`
	// Attachments any attached files
	Attachments []Attachment `json:"attachments"`
	// Embeds any embedded content
	Embeds []Embed `json:"embeds"`
	// Reactions reactions to the message
	Reactions []Reaction `json:"reactions"`
	// Nonce used for validating a message was sent
	Nonce string `json:"nonce"`
	// Pinned whether this message is pinned
	Pinned bool `json:"pinned"`
	// WebhookId if the message is generated by a webhook, this is the webhook's id
	WebhookId string `json:"webhook_id"`
	// Type [type of message](#DOCS_RESOURCES_CHANNEL/message-object-message-types)
	Type int `json:"type"`
	// Activity sent with Rich Presence-related chat embeds
	Activity MessageActivity `json:"activity"`
	// Application sent with Rich Presence-related chat embeds
	Application Application `json:"application"`
	// ApplicationId if the message is a response to an [Interaction](#DOCS_INTERACTIONS_SLASH_COMMANDS/), this is the id of the interaction's application
	ApplicationId string `json:"application_id"`
	// MessageReference data showing the source of a crosspost, channel follow add, pin, or reply message
	MessageReference MessageReference `json:"message_reference"`
	// Flags [message flags](#DOCS_RESOURCES_CHANNEL/message-object-message-flags) combined as a [bitfield](https://en.wikipedia.org/wiki/Bit_field)
	Flags int `json:"flags"`
	// ReferencedMessage the message associated with the message_reference
	ReferencedMessage *Message `json:"referenced_message"`
	// Interaction sent if the message is a response to an [Interaction](#DOCS_INTERACTIONS_SLASH_COMMANDS/)
	Interaction InteractionResponse `json:"interaction"`
	// Thread the thread that was started from this message, includes [thread member](#DOCS_RESOURCES_CHANNEL/thread-member-object) object
	Thread Channel `json:"thread"`
	// Components sent if the message contains components like buttons, action rows, or other interactive components
	Components []Component `json:"components"`
	// StickerItems sent if the message contains stickers
	StickerItems []StickerItem `json:"sticker_items"`
	// Stickers **Deprecated** the stickers sent with the message
	Stickers []Sticker `json:"stickers"`
}

type MessageActivity struct {
	// Type [type of message activity](#DOCS_RESOURCES_CHANNEL/message-object-message-activity-types)
	Type int `json:"type"`
	// PartyId party_id from a [Rich Presence event](#DOCS_RICH_PRESENCE_HOW_TO/updating-presence-update-presence-payload-fields)
	PartyId string `json:"party_id"`
}

type MessageReference struct {
	// MessageId id of the originating message
	MessageId string `json:"message_id"`
	// ChannelId id of the originating message's channel
	ChannelId string `json:"channel_id"`
	// GuildId id of the originating message's guild
	GuildId string `json:"guild_id"`
	// FailIfNotExists when sending, whether to error if the referenced message doesn't exist instead of sending as a normal (non-reply) message, default true
	FailIfNotExists bool `json:"fail_if_not_exists"`
}

type Overwrite struct {
	// Id role or user id
	Id string `json:"id"`
	// Type either 0 (role) or 1 (member)
	Type int `json:"type"`
	// Allow permission bit set
	Allow string `json:"allow"`
	// Deny permission bit set
	Deny string `json:"deny"`
}

type Reaction struct {
	// Count times this emoji has been used to react
	Count int `json:"count"`
	// Me whether the current user reacted using this emoji
	Me bool `json:"me"`
	// Emoji emoji information
	Emoji Emoji `json:"emoji"`
}

type Role struct {
	// Id role id
	Id string `json:"id"`
	// Name role name
	Name string `json:"name"`
	// Color integer representation of hexadecimal color code
	Color int `json:"color"`
	// Hoist if this role is pinned in the user listing
	Hoist bool `json:"hoist"`
	// Position position of this role
	Position int `json:"position"`
	// Permissions permission bit set
	Permissions string `json:"permissions"`
	// Managed whether this role is managed by an integration
	Managed bool `json:"managed"`
	// Mentionable whether this role is mentionable
	Mentionable bool `json:"mentionable"`
	// Tags the tags this role has
	Tags RoleTags `json:"tags"`
}

type RoleTags struct {
	// BotId the id of the bot this role belongs to
	BotId string `json:"bot_id"`
	// IntegrationId the id of the integration this role belongs to
	IntegrationId string `json:"integration_id"`
	// PremiumSubscriber whether this is the guild's premium subscriber role
	PremiumSubscriber interface{} `json:"premium_subscriber"`
}

type SelectOption struct {
	// Label the user-facing name of the option, max 25 characters
	Label string `json:"label"`
	// Value the dev-define value of the option, max 100 characters
	Value string `json:"value"`
	// Description an additional description of the option, max 50 characters
	Description string `json:"description"`
	// Emoji `id`, `name`, and `animated`
	Emoji Emoji `json:"emoji"`
	// Default will render this option as selected by default
	Default bool `json:"default"`
}

type StageInstance struct {
	// Id The id of this Stage instance
	Id string `json:"id"`
	// GuildId The guild id of the associated Stage channel
	GuildId string `json:"guild_id"`
	// ChannelId The id of the associated Stage channel
	ChannelId string `json:"channel_id"`
	// Topic The topic of the Stage instance (1-120 characters)
	Topic string `json:"topic"`
	// PrivacyLevel The [privacy level](#DOCS_RESOURCES_STAGE_INSTANCE/stage-instance-object-privacy-level) of the Stage instance
	PrivacyLevel int `json:"privacy_level"`
	// DiscoverableDisabled Whether or not Stage Discovery is disabled
	DiscoverableDisabled bool `json:"discoverable_disabled"`
}

type Sticker struct {
	// Id [id of the sticker](#DOCS_REFERENCE/image-formatting)
	Id string `json:"id"`
	// PackId for standard stickers, id of the pack the sticker is from
	PackId string `json:"pack_id"`
	// Name name of the sticker
	Name string `json:"name"`
	// Description description of the sticker
	Description string `json:"description"`
	// Tags for guild stickers, the Discord name of a unicode emoji representing the sticker's expression. for standard stickers, a comma-separated list of related expressions.
	Tags string `json:"tags"`
	// Asset **Deprecated** previously the sticker asset hash, now an empty string
	Asset string `json:"asset"`
	// Type [type of sticker](#DOCS_RESOURCES_STICKER/sticker-object-sticker-types)
	Type int `json:"type"`
	// FormatType [type of sticker format](#DOCS_RESOURCES_STICKER/sticker-object-sticker-format-types)
	FormatType int `json:"format_type"`
	// Available whether this guild sticker can be used, may be false due to loss of Server Boosts
	Available bool `json:"available"`
	// GuildId id of the guild that owns this sticker
	GuildId string `json:"guild_id"`
	// User the user that uploaded the guild sticker
	User User `json:"user"`
	// SortValue the standard sticker's sort order within its pack
	SortValue int `json:"sort_value"`
}

type StickerItem struct {
	// Id id of the sticker
	Id string `json:"id"`
	// Name name of the sticker
	Name string `json:"name"`
	// FormatType [type of sticker format](#DOCS_RESOURCES_STICKER/sticker-object-sticker-format-types)
	FormatType int `json:"format_type"`
}
type Team struct {
	// Icon a hash of the image of the team's icon
	Icon string `json:"icon"`
	// Id the unique id of the team
	Id string `json:"id"`
	// Members the members of the team
	Members []TeamMember `json:"members"`
	// Name the name of the team
	Name string `json:"name"`
	// OwnerUserId the user id of the current team owner
	OwnerUserId string `json:"owner_user_id"`
}

type TeamMember struct {
	// MembershipState the user's [membership state](#DOCS_TOPICS_TEAMS/data-models-membership-state-enum) on the team
	MembershipState int `json:"membership_state"`
	// Permissions will always be `["*"]`
	Permissions []string `json:"permissions"`
	// TeamId the id of the parent team of which they are a member
	TeamId string `json:"team_id"`
	// User the avatar, discriminator, id, and username of the user
	User User `json:"user"`
}

type ThreadMetadata struct {
	// Archived whether the thread is archived
	Archived bool `json:"archived"`
	// AutoArchiveDuration duration in minutes to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	AutoArchiveDuration int `json:"auto_archive_duration"`
	// ArchiveTimestamp timestamp when the thread's archive status was last changed, used for calculating recent activity
	ArchiveTimestamp time.Time `json:"archive_timestamp"`
	// Locked whether the thread is locked; when a thread is locked, only users with MANAGE_THREADS can unarchive it
	Locked bool `json:"locked"`
}

type ThreadMember struct {
	// Id the id of the thread
	Id string `json:"id"`
	// UserId the id of the user
	UserId string `json:"user_id"`
	// JoinTimestamp the time the current user last joined the thread
	JoinTimestamp time.Time `json:"join_timestamp"`
	// Flags any user-thread settings, currently only used for notifications
	Flags int `json:"flags"`
}

type User struct {
	// Id the user's id
	Id string `json:"id"`
	// Username the user's username, not unique across the platform
	Username string `json:"username"`
	// Discriminator the user's 4-digit discord-tag
	Discriminator string `json:"discriminator"`
	// Avatar the user's [avatar hash](#DOCS_REFERENCE/image-formatting)
	Avatar string `json:"avatar"`
	// Bot whether the user belongs to an OAuth2 application
	Bot bool `json:"bot"`
	// System whether the user is an Official Discord System user (part of the urgent message system)
	System bool `json:"system"`
	// MfaEnabled whether the user has two factor enabled on their account
	MfaEnabled bool `json:"mfa_enabled"`
	// Locale the user's chosen language option
	Locale string `json:"locale"`
	// Verified whether the email on this account has been verified
	Verified bool `json:"verified"`
	// Email the user's email
	Email string `json:"email"`
	// Flags the [flags](#DOCS_RESOURCES_USER/user-object-user-flags) on a user's account
	Flags int `json:"flags"`
	// PremiumType the [type of Nitro subscription](#DOCS_RESOURCES_USER/user-object-premium-types) on a user's account
	PremiumType int `json:"premium_type"`
	// PublicFlags the public [flags](#DOCS_RESOURCES_USER/user-object-user-flags) on a user's account
	PublicFlags int `json:"public_flags"`
	// Member only included with partial data for some guild based purposes
	Member *GuildMember `json:"member,omitempty"`
}

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
