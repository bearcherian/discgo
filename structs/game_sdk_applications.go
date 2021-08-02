package structs

type OAuth2Token struct {
	// AccessToken a bearer token for the current user
	AccessToken string `json:"AccessToken"`
	// Scopes a list of oauth2 scopes as a single string, delineated by spaces like `"identify rpc gdm.join"`
	Scopes string `json:"Scopes"`
	// Expires the timestamp at which the token expires
	Expires int64 `json:"Expires"`
}

type SignedAppTicket struct {
	// ApplicationId the application id for the ticket
	ApplicationId int64 `json:"application_id"`
	// User the user for the ticket
	User User `json:"user"`
	// Entitlements the list of the user's entitlements for this application
	Entitlements []Entitlement `json:"entitlements"`
	// Timestamp the ISO 8601 timestamp for the ticket
	Timestamp string `json:"timestamp"`
}
