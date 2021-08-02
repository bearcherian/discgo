package structs

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
	PremiumSubscriber null `json:"premium_subscriber"`
}