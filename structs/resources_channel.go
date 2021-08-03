package structs

import "time"

type AllowedMentionType string

const (
	AllowedMentionRole     AllowedMentionType = "roles"
	AllowedMentionUser     AllowedMentionType = "users"
	AllowedMentionEveryone AllowedMentionType = "everyone"
)

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

type FollowedChannel struct {
	// ChannelId source channel id
	ChannelId string `json:"channel_id"`
	// WebhookId created target webhook id
	WebhookId string `json:"webhook_id"`
}

type Reaction struct {
	// Count times this emoji has been used to react
	Count int `json:"count"`
	// Me whether the current user reacted using this emoji
	Me bool `json:"me"`
	// Emoji emoji information
	Emoji Emoji `json:"emoji"`
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
