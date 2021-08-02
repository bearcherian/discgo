package structs

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