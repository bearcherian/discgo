package structs

type ApplicationCommandPermissionType int
type ApplicationCommandOptionType int
type InteractionRequestType int
type InteractionCallbackType int

const (
	ApplicationCommandPermissionRole ApplicationCommandPermissionType = 1
	ApplicationCommandPermissionUser ApplicationCommandPermissionType = 2

	ApplicationCommandOptionSubCommand      ApplicationCommandOptionType = 1
	ApplicationCommandOptionSubCommandGroup ApplicationCommandOptionType = 2
	ApplicationCommandOptionString          ApplicationCommandOptionType = 3
	ApplicationCommandOptionInteger         ApplicationCommandOptionType = 4
	ApplicationCommandOptionBoolean         ApplicationCommandOptionType = 5
	ApplicationCommandOptionUser            ApplicationCommandOptionType = 6
	ApplicationCommandOptionChannel         ApplicationCommandOptionType = 7
	ApplicationCommandOptionRole            ApplicationCommandOptionType = 8
	ApplicationCommandOptionMentionable     ApplicationCommandOptionType = 9
	ApplicationCommandOptionNumber          ApplicationCommandOptionType = 10

	InteractionRequestPing               InteractionRequestType = 1
	InteractionRequestApplicationCommand InteractionRequestType = 2
	InteractionRequestMessageComponent   InteractionRequestType = 3

	InteractionCallbackPong                             InteractionCallbackType = 1
	InteractionCallbackChannelMessageWithSource         InteractionCallbackType = 4
	InteractionCallbackDeferredChannelMessageWithSource InteractionCallbackType = 5
	InteractionCallbackDeferredUpdateMessage            InteractionCallbackType = 6
	InteractionCallbackUpdateMessage                    InteractionCallbackType = 7
)

type ApplicationCommand struct {
	// Id unique id of the command
	Id string `json:"id"`
	// ApplicationId unique id of the parent application
	ApplicationId string `json:"application_id"`
	// GuildId guild id of the command, if not global
	GuildId string `json:"guild_id"`
	// Name 1-32 lowercase character name matching `^[\w-]{1,32}$`
	Name string `json:"name"`
	// Description 1-100 character description
	Description string `json:"description"`
	// Options the parameters for the command
	Options []ApplicationCommandOption `json:"options"`
	// DefaultPermission whether the command is enabled by default when the app is added to a guild. (default true)
	DefaultPermission bool `json:"default_permission"`
}

type ApplicationCommandOption struct {
	// Type value of [application command option type](#DOCS_INTERACTIONS_SLASH_COMMANDS/application-command-object-application-command-option-type)
	Type int `json:"type"`
	// Name 1-32 lowercase character name matching `^[\w-]{1,32}$`
	Name string `json:"name"`
	// Description 1-100 character description
	Description string `json:"description"`
	// Required if the parameter is required or optional--default `false`
	Required bool `json:"required"`
	// Choices choices for `STRING`, `INTEGER`, and `NUMBER` types for the user to pick from
	Choices []ApplicationCommandOptionChoice `json:"choices"`
	// Options if the option is a subcommand or subcommand group type, this nested options will be the parameters
	Options []ApplicationCommandOption `json:"options"`
}

type ApplicationCommandOptionChoice struct {
	// Name 1-100 character choice name
	Name string `json:"name"`
	// Value value of the choice, up to 100 characters if string, string, integer or double
	Value interface{} `json:"value"`
}

type GuildApplicationCommandPermissions struct {
	// Id the id of the command
	Id string `json:"id"`
	// ApplicationId the id of the application the command belongs to
	ApplicationId string `json:"application_id"`
	// GuildId the id of the guild
	GuildId string `json:"guild_id"`
	// Permissions the permissions for the command in the guild
	Permissions []GuildApplicationCommandPermissions `json:"permissions"`
}

type ApplicationCommandPermissions struct {
	// Id the id of the role or user
	Id string `json:"id"`
	// Type role or user
	Type ApplicationCommandPermissionType `json:"type"`
	// Permission `true` to allow, `false`, to disallow
	Permission bool `json:"permission"`
}

type Interaction struct {
	// Id id of the interaction
	Id string `json:"id"`
	// ApplicationId id of the application this interaction is for
	ApplicationId string `json:"application_id"`
	// Type the type of interaction
	Type InteractionRequestType `json:"type"`
	// Data the command data payload
	Data ApplicationCommandInteractionData `json:"data"`
	// GuildId the guild it was sent from
	GuildId string `json:"guild_id"`
	// ChannelId the channel it was sent from
	ChannelId string `json:"channel_id"`
	// Member guild member data for the invoking user, including permissions
	Member GuildMember `json:"member"`
	// User user object for the invoking user, if invoked in a DM
	User User `json:"user"`
	// Token a continuation token for responding to the interaction
	Token string `json:"token"`
	// Version read-only property, always `1`
	Version int `json:"version"`
	// Message for components, the message they were attached to
	Message Message `json:"message"`
}

type ApplicationCommandInteractionData struct {
	// Id the ID of the invoked command
	Id string `json:"id"`
	// Name the name of the invoked command
	Name string `json:"name"`
	// Resolved converted users + roles + channels
	Resolved ApplicationCommandInteractionDataResolved `json:"resolved"`
	// Options the params + values from the user
	Options []ApplicationCommandInteractionDataOption `json:"options"`
	// CustomId for components, the [`custom_id`](#DOCS_INTERACTIONS_MESSAGE_COMPONENTS/custom-id) of the component
	CustomId string `json:"custom_id"`
	// ComponentType for components, the [type](#DOCS_INTERACTIONS_MESSAGE_COMPONENTS/component-types) of the component
	ComponentType int `json:"component_type"`
}

type ApplicationCommandInteractionDataResolved struct {
	// Users the ids and User objects
	Users User `json:"users"`
	// Members the ids and partial Member objects
	Members TeamMember `json:"members"`
	// Roles the ids and Role objects
	Roles Role `json:"roles"`
	// Channels the ids and partial Channel objects
	Channels Channel `json:"channels"`
}

type ApplicationCommandInteractionDataOption struct {
	// Name the name of the parameter
	Name string `json:"name"`
	// Type value of [application command option type](#DOCS_INTERACTIONS_SLASH_COMMANDS/application-command-object-application-command-option-type)
	Type int `json:"type"`
	// Value the value of the pair
	Value ApplicationCommandOptionType `json:"value"`
	// Options present if this option is a group or subcommand
	Options []ApplicationCommandInteractionDataOption `json:"options"`
}

type InteractionResponse struct {
	// Type the type of response
	Type InteractionCallbackType `json:"type"`
	// Data an optional response message
	Data InteractionApplicationCommandCallbackData `json:"data"`
}

type InteractionApplicationCommandCallbackData struct {
	// Tts is the response TTS
	Tts bool `json:"tts"`
	// Content message content
	Content string `json:"content"`
	// Embeds supports up to 10 embeds
	Embeds []Embed `json:"embeds"`
	// AllowedMentions [allowed mentions](#DOCS_RESOURCES_CHANNEL/allowed-mentions-object) object
	AllowedMentions AllowedMentions `json:"allowed_mentions"`
	// Flags [interaction application command callback data flags](#DOCS_INTERACTIONS_SLASH_COMMANDS/interaction-response-object-interaction-application-command-callback-data-flags)
	Flags int `json:"flags"`
	// Components message components
	Components []Component `json:"components"`
}

type MessageInteraction struct {
	// Id id of the interaction
	Id string `json:"id"`
	// Type the type of interaction
	Type InteractionRequestType `json:"type"`
	// Name the name of the [application command](#DOCS_INTERACTIONS_SLASH_COMMANDS/application-command-object-application-command-structure)
	Name string `json:"name"`
	// User the user who invoked the interaction
	User User `json:"user"`
}
