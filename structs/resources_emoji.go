package structs

type Emoji struct {
	// Id [emoji id](#DOCS_REFERENCE/image-formatting)
	Id string `json:"id"`
	// Name emoji name
	Name string (can be null only in reaction emoji objects) `json:"name"`
	// Roles roles allowed to use this emoji
	Roles array of [role](#DOCS_TOPICS_PERMISSIONS/role-object) object ids `json:"roles"`
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