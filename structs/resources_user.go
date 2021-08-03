package structs

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
	Integrations []Integration `json:"integrations"`
	// Verified whether the connection is verified
	Verified bool `json:"verified"`
	// FriendSync whether friend sync is enabled for this connection
	FriendSync bool `json:"friend_sync"`
	// ShowActivity whether activities related to this connection will be shown in presence updates
	ShowActivity bool `json:"show_activity"`
	// Visibility [visibility](#DOCS_RESOURCES_USER/connection-object-visibility-types) of this connection
	Visibility int `json:"visibility"`
}
