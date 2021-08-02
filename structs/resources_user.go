package structs

type Connection struct {
	// Id id of the connection account
	Id string `json:"id"`
	// Name the username of the connection account
	Name string `json:"name"`
	// Type the service of the connection (twitch, youtube)
	Type string `json:"type"`
	// Revoked whether the connection is revoked
	Revoked bool `json:"revoked"`
	// Integrations an array of partial [server integrations](#DOCS_RESOURCES_GUILD/integration-object)
	Integrations array `json:"integrations"`
	// Verified whether the connection is verified
	Verified bool `json:"verified"`
	// FriendSync whether friend sync is enabled for this connection
	FriendSync bool `json:"friend_sync"`
	// ShowActivity whether activities related to this connection will be shown in presence updates
	ShowActivity bool `json:"show_activity"`
	// Visibility [visibility](#DOCS_RESOURCES_USER/connection-object-visibility-types) of this connection
	Visibility int `json:"visibility"`
}