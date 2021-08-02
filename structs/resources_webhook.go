package structs

type Webhook struct {
	// Id the id of the webhook
	Id string `json:"id"`
	// Type the [type](#DOCS_RESOURCES_WEBHOOK/webhook-object-webhook-types) of the webhook
	Type int `json:"type"`
	// GuildId the guild id this webhook is for, if any
	GuildId string `json:"guild_id"`
	// ChannelId the channel id this webhook is for, if any
	ChannelId string `json:"channel_id"`
	// User the user this webhook was created by (not returned when getting a webhook with its token)
	User User `json:"user"`
	// Name the default name of the webhook
	Name string `json:"name"`
	// Avatar the default user avatar [hash](#DOCS_REFERENCE/image-formatting) of the webhook
	Avatar string `json:"avatar"`
	// Token the secure token of the webhook (returned for Incoming Webhooks)
	Token string `json:"token"`
	// ApplicationId the bot/OAuth2 application that created this webhook
	ApplicationId string `json:"application_id"`
	// SourceGuild the guild of the channel that this webhook is following (returned for Channel Follower Webhooks)
	SourceGuild Guild `json:"source_guild"`
	// SourceChannel the channel that this webhook is following (returned for Channel Follower Webhooks)
	SourceChannel Channel `json:"source_channel"`
	// Url the url used for executing the webhook (returned by the [webhooks](#DOCS_TOPICS_OAUTH2/webhooks) OAuth2 flow)
	Url string `json:"url"`
}