package structs

type GuildTemplate struct {
	// Code the template code (unique ID)
	Code string `json:"code"`
	// Name template name
	Name string `json:"name"`
	// Description the description for the template
	Description string `json:"description"`
	// UsageCount number of times this template has been used
	UsageCount int `json:"usage_count"`
	// CreatorId the ID of the user who created the template
	CreatorId string `json:"creator_id"`
	// Creator the user who created the template
	Creator User `json:"creator"`
	// CreatedAt when this template was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt when this template was last synced to the source guild
	UpdatedAt time.Time `json:"updated_at"`
	// SourceGuildId the ID of the guild this template is based on
	SourceGuildId string `json:"source_guild_id"`
	// SerializedSourceGuild the guild snapshot this template contains
	SerializedSourceGuild Guild `json:"serialized_source_guild"`
	// IsDirty whether the template has unsynced changes
	IsDirty bool `json:"is_dirty"`
}