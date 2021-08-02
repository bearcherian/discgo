package structs

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