package structs

type Component struct {
	// Type [component type](#DOCS_INTERACTIONS_MESSAGE_COMPONENTS/component-object-component-types)
	Type int `json:"type"`
	// CustomId a developer-defined identifier for the component, max 100 characters
	CustomId string `json:"custom_id"`
	// Disabled whether the component is disabled, default `false`
	Disabled bool `json:"disabled"`
	// Style one of [button styles](#DOCS_INTERACTIONS_MESSAGE_COMPONENTS/button-object-button-styles)
	Style int `json:"style"`
	// Label text that appears on the button, max 80 characters
	Label string `json:"label"`
	// Emoji `name`, `id`, and `animated`
	Emoji Emoji `json:"emoji"`
	// Url a url for link-style buttons
	Url string `json:"url"`
	// Options the choices in the select, max 25
	Options []SelectOption `json:"options"`
	// Placeholder custom placeholder text if nothing is selected, max 100 characters
	Placeholder string `json:"placeholder"`
	// MinValues the minimum number of items that must be chosen; default 1, min 0, max 25
	MinValues int `json:"min_values"`
	// MaxValues the maximum number of items that can be chosen; default 1, max 25
	MaxValues int `json:"max_values"`
	// Components a list of child components
	Components []Component `json:"components"`
}

type Button struct {
	// Type `2` for a button
	Type int `json:"type"`
	// Style one of [button styles](#DOCS_INTERACTIONS_MESSAGE_COMPONENTS/buttons-button-styles)
	Style int `json:"style"`
	// Label text that appears on the button, max 80 characters
	Label string `json:"label"`
	// Emoji `name`, `id`, and `animated`
	Emoji Emoji `json:"emoji"`
	// CustomId a developer-defined identifier for the button, max 100 characters
	CustomId string `json:"custom_id"`
	// Url a url for link-style buttons
	Url string `json:"url"`
	// Disabled whether the button is disabled (default `false`)
	Disabled bool `json:"disabled"`
}

type SelectMenu struct {
	// Type `3` for a select menu
	Type int `json:"type"`
	// CustomId a developer-defined identifier for the button, max 100 characters
	CustomId string `json:"custom_id"`
	// Options the choices in the select, max 25
	Options []SelectOption `json:"options"`
	// Placeholder custom placeholder text if nothing is selected, max 100 characters
	Placeholder string `json:"placeholder"`
	// MinValues the minimum number of items that must be chosen; default 1, min 0, max 25
	MinValues int `json:"min_values"`
	// MaxValues the maximum number of items that can be chosen; default 1, max 25
	MaxValues int `json:"max_values"`
	// Disabled disable the select, default false
	Disabled bool `json:"disabled"`
}

type SelectOption struct {
	// Label the user-facing name of the option, max 25 characters
	Label string `json:"label"`
	// Value the dev-define value of the option, max 100 characters
	Value string `json:"value"`
	// Description an additional description of the option, max 50 characters
	Description string `json:"description"`
	// Emoji `id`, `name`, and `animated`
	Emoji Emoji `json:"emoji"`
	// Default will render this option as selected by default
	Default bool `json:"default"`
}
