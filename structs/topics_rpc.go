package structs

type Payload struct {
	// Cmd [payload command](#DOCS_TOPICS_RPC/commands-and-events-rpc-commands)
	Cmd enum `json:"cmd"`
	// Nonce unique string used once for replies from the server
	Nonce string `json:"nonce"`
	// Evt [subscription event](#DOCS_TOPICS_RPC/commands-and-events-rpc-events)
	Evt enum `json:"evt"`
	// Data event data
	Data  `json:"data"`
	// Args command arguments
	Args  `json:"args"`
}

type AuthorizeArgument struct {
	// Scopes scopes to authorize
	Scopes array of [OAuth2 scopes](#DOCS_TOPICS_OAUTH2/scopes) `json:"scopes"`
	// ClientId OAuth2 application id
	ClientId string `json:"client_id"`
	// RpcToken one-time use RPC token
	RpcToken string `json:"rpc_token"`
	// Username username to create a guest account with if the user does not have Discord
	Username string `json:"username"`
}

type AuthorizeResponse struct {
	// Code OAuth2 authorization code
	Code string `json:"code"`
}

type AuthenticateArgument struct {
	// AccessToken OAuth2 access token
	AccessToken string `json:"access_token"`
}

type AuthenticateResponse struct {
	// User the authed user
	User User `json:"user"`
	// Scopes authorized scopes
	Scopes array of [OAuth2 scopes](#DOCS_TOPICS_OAUTH2/scopes) `json:"scopes"`
	// Expires expiration date of OAuth2 token
	Expires date `json:"expires"`
	// Application application the user authorized
	Application OAuth2Application `json:"application"`
}

type OAuth2Application struct {
	// Description application description
	Description string `json:"description"`
	// Icon hash of the icon
	Icon string `json:"icon"`
	// Id application client id
	Id string `json:"id"`
	// RpcOrigins array of rpc origin urls
	RpcOrigins []string `json:"rpc_origins"`
	// Name application name
	Name string `json:"name"`
}

type GetGuildsResponse struct {
	// Guilds the guilds the user is in
	Guilds []Guild `json:"guilds"`
}

type GetGuildArgument struct {
	// GuildId id of the guild to get
	GuildId string `json:"guild_id"`
	// Timeout asynchronously get guild with time to wait before timing out
	Timeout int `json:"timeout"`
}

type GetGuildResponse struct {
	// Id guild id
	Id string `json:"id"`
	// Name guild name
	Name string `json:"name"`
	// IconUrl guild icon url
	IconUrl string `json:"icon_url"`
	// Members members of the guild (deprecated; always empty array)
	Members []GuildMember `json:"members"`
}

type GetChannelArgument struct {
	// ChannelId id of the channel to get
	ChannelId string `json:"channel_id"`
}

type GetChannelResponse struct {
	// Id channel id
	Id string `json:"id"`
	// GuildId channel's guild id
	GuildId string `json:"guild_id"`
	// Name channel name
	Name string `json:"name"`
	// Type channel type (guild text: 0, guild voice: 2, dm: 1, group dm: 3)
	Type int `json:"type"`
	// Topic (text) channel topic
	Topic string `json:"topic"`
	// Bitrate (voice) bitrate of voice channel
	Bitrate int `json:"bitrate"`
	// UserLimit (voice) user limit of voice channel (0 for none)
	UserLimit int `json:"user_limit"`
	// Position position of channel in channel list
	Position int `json:"position"`
	// VoiceStates (voice) channel's voice states
	VoiceStates []VoiceState `json:"voice_states"`
	// Messages (text) channel's messages
	Messages []Message `json:"messages"`
}

type GetChannelsArgument struct {
	// GuildId id of the guild to get channels for
	GuildId string `json:"guild_id"`
}

type GetChannelsResponse struct {
	// Channels guild channels the user is in
	Channels []Channel `json:"channels"`
}

type SetUserVoiceSettingsArgumentandResponse struct {
	// UserId user id
	UserId string `json:"user_id"`
	// Pan set the pan of the user
	Pan Pan `json:"pan"`
	// Volume set the volume of user (defaults to 100, min 0, max 200)
	Volume int `json:"volume"`
	// Mute set the mute state of the user
	Mute bool `json:"mute"`
}

type Pan struct {
	// Left left pan of user (min: 0.0, max: 1.0)
	Left float32 `json:"left"`
	// Right right pan of user (min: 0.0, max: 1.0)
	Right float32 `json:"right"`
}

type SelectVoiceChannelArgument struct {
	// ChannelId channel id to join (or `null` to leave)
	ChannelId string `json:"channel_id"`
	// Timeout asynchronously join channel with time to wait before timing out
	Timeout int `json:"timeout"`
	// Force forces a user to join a voice channel
	Force bool `json:"force"`
}

type SelectTextChannelArgument struct {
	// ChannelId channel id to join (or `null` to leave)
	ChannelId string `json:"channel_id"`
	// Timeout asynchronously join channel with time to wait before timing out
	Timeout int `json:"timeout"`
}

type GetVoiceSettingsResponse struct {
	// Input input settings
	Input VoiceSettingsInput `json:"input"`
	// Output output settings
	Output VoiceSettingsOutput `json:"output"`
	// Mode voice mode settings
	Mode VoiceSettingsMode `json:"mode"`
	// AutomaticGainControl state of automatic gain control
	AutomaticGainControl bool `json:"automatic_gain_control"`
	// EchoCancellation state of echo cancellation
	EchoCancellation bool `json:"echo_cancellation"`
	// NoiseSuppression state of noise suppression
	NoiseSuppression bool `json:"noise_suppression"`
	// Qos state of voice quality of service
	Qos bool `json:"qos"`
	// SilenceWarning state of silence warning notice
	SilenceWarning bool `json:"silence_warning"`
	// Deaf state of self-deafen
	Deaf bool `json:"deaf"`
	// Mute state of self-mute
	Mute bool `json:"mute"`
}

type VoiceSettingsInput struct {
	// DeviceId device id
	DeviceId string `json:"device_id"`
	// Volume input voice level (min: 0, max: 100)
	Volume float32 `json:"volume"`
	// AvailableDevices array of _read-only_ device objects containing `id` and `name` string keys
	AvailableDevices [] `json:"available_devices"`
}

type VoiceSettingsOutput struct {
	// DeviceId device id
	DeviceId string `json:"device_id"`
	// Volume output voice level (min: 0, max: 200)
	Volume float32 `json:"volume"`
	// AvailableDevices array of _read-only_ device objects containing `id` and `name` string keys
	AvailableDevices [] `json:"available_devices"`
}

type VoiceSettingsMode struct {
	// Type voice setting mode type (can be `PUSH_TO_TALK` or `VOICE_ACTIVITY`)
	Type string `json:"type"`
	// AutoThreshold voice activity threshold automatically sets its threshold
	AutoThreshold bool `json:"auto_threshold"`
	// Threshold threshold for voice activity (in dB) (min: -100, max: 0)
	Threshold float32 `json:"threshold"`
	// Shortcut shortcut key combos for PTT
	Shortcut ShortcutKeyCombo `json:"shortcut"`
	// Delay the PTT release delay (in ms) (min: 0, max: 2000)
	Delay float32 `json:"delay"`
}

type ShortcutKeyCombo struct {
	// Type see [key types](#DOCS_TOPICS_RPC/shortcut-key-combo-object-key-types)
	Type int `json:"type"`
	// Code key code
	Code int `json:"code"`
	// Name key name
	Name string `json:"name"`
}

type SetVoiceSettingsArgumentandResponse struct {
	// Input input settings
	Input VoiceSettingsInput `json:"input"`
	// Output output settings
	Output VoiceSettingsOutput `json:"output"`
	// Mode voice mode settings
	Mode VoiceSettingsMode `json:"mode"`
	// AutomaticGainControl state of automatic gain control
	AutomaticGainControl bool `json:"automatic_gain_control"`
	// EchoCancellation state of echo cancellation
	EchoCancellation bool `json:"echo_cancellation"`
	// NoiseSuppression state of noise suppression
	NoiseSuppression bool `json:"noise_suppression"`
	// Qos state of voice quality of service
	Qos bool `json:"qos"`
	// SilenceWarning state of silence warning notice
	SilenceWarning bool `json:"silence_warning"`
	// Deaf state of self-deafen
	Deaf bool `json:"deaf"`
	// Mute state of self-mute
	Mute bool `json:"mute"`
}

type SubscribeResponse struct {
	// Evt event name now subscribed to
	Evt string `json:"evt"`
}

type UnsubscribeResponse struct {
	// Evt event name now unsubscribed from
	Evt string `json:"evt"`
}

type SetActivityArgument struct {
	// Pid the application's process id
	Pid int `json:"pid"`
	// Activity the rich presence to assign to the user
	Activity Activity `json:"activity"`
}

type SendActivityJoinInviteArgument struct {
	// UserId the id of the requesting user
	UserId string `json:"user_id"`
}

type CloseActivityRequestArgument struct {
	// UserId the id of the requesting user
	UserId string `json:"user_id"`
}

type ReadyDispatchData struct {
	// V RPC version
	V int `json:"v"`
	// Config server configuration
	Config RpcServerConfiguration `json:"config"`
	// User the user to whom you are connected
	User User `json:"user"`
}

type RPCServerConfiguration struct {
	// CdnHost server's cdn
	CdnHost string `json:"cdn_host"`
	// ApiEndpoint server's api endpoint
	ApiEndpoint string `json:"api_endpoint"`
	// Environment server's environment
	Environment string `json:"environment"`
}

type ErrorData struct {
	// Code RPC Error Code
	Code int `json:"code"`
	// Message Error description
	Message string `json:"message"`
}

type GuildStatusArgument struct {
	// GuildId id of guild to listen to updates of
	GuildId string `json:"guild_id"`
}

type GuildStatusDispatchData struct {
	// Guild guild with requested id
	Guild Guild `json:"guild"`
	// Online number of online users in guild (deprecated; always 0)
	Online int `json:"online"`
}

type GuildCreateDispatchData struct {
	// Id guild id
	Id string `json:"id"`
	// Name name of the guild
	Name string `json:"name"`
}

type ChannelCreateDispatchData struct {
	// Id channel id
	Id string `json:"id"`
	// Name name of the channel
	Name string `json:"name"`
	// Type channel type (guild text: 0, guild voice: 2, dm: 1, group dm: 3)
	Type int `json:"type"`
}

type VoiceChannelSelectDispatchData struct {
	// ChannelId id of channel (`null` if none)
	ChannelId string `json:"channel_id"`
	// GuildId id of guild (`null` if none)
	GuildId string `json:"guild_id"`
}

type VoiceStateArgument struct {
	// ChannelId id of channel to listen to updates of
	ChannelId string `json:"channel_id"`
}

type VoiceConnectionStatusDispatchData struct {
	// State one of the voice connection states listed below
	State string `json:"state"`
	// Hostname hostname of the connected voice server
	Hostname string `json:"hostname"`
	// Pings last 20 pings (in ms)
	Pings array of integers `json:"pings"`
	// AveragePing average ping (in ms)
	AveragePing int `json:"average_ping"`
	// LastPing last ping (in ms)
	LastPing int `json:"last_ping"`
}

type MessageArgument struct {
	// ChannelId id of channel to listen to updates of
	ChannelId string `json:"channel_id"`
}

type SpeakingArgument struct {
	// ChannelId id of channel to listen to updates of
	ChannelId string `json:"channel_id"`
}

type SpeakingDispatchData struct {
	// UserId id of user who started/stopped speaking
	UserId string `json:"user_id"`
}

type NotificationCreateDispatchData struct {
	// ChannelId id of channel where notification occurred
	ChannelId string `json:"channel_id"`
	// Message message that generated this notification
	Message Message `json:"message"`
	// IconUrl icon url of the notification
	IconUrl string `json:"icon_url"`
	// Title title of the notification
	Title string `json:"title"`
	// Body body of the notification
	Body string `json:"body"`
}

type ActivityJoinDispatchData struct {
	// Secret the [`join_secret`](#DOCS_RICH_PRESENCE_HOW_TO/updating-presence-update-presence-payload-fields) for the given invite
	Secret string `json:"secret"`
}

type ActivitySpectateDispatchData struct {
	// Secret the [`spectate_secret`](#DOCS_RICH_PRESENCE_HOW_TO/updating-presence-update-presence-payload-fields) for the given invite
	Secret string `json:"secret"`
}

type ActivityJoinRequestData struct {
	// User information about the user requesting to join
	User User `json:"user"`
}