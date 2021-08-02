package structs

type User struct {
	// Id the user's id
	Id int64 `json:"Id"`
	// Username their name
	Username string `json:"Username"`
	// Discriminator the user's unique discrim
	Discriminator string `json:"Discriminator"`
	// Avatar the hash of the user's avatar
	Avatar string `json:"Avatar"`
	// Bot if the user is a bot user
	Bot bool `json:"Bot"`
}

type Activity struct {
	// ApplicationId your application id - this is a read-only field
	ApplicationId int64 `json:"ApplicationId"`
	// Name name of the application - this is a read-only field
	Name string `json:"Name"`
	// State the player's current party status
	State string `json:"State"`
	// Details what the player is currently doing
	Details string `json:"Details"`
	// Timestamps helps create elapsed/remaining timestamps on a player's profile
	Timestamps ActivityTimestamps `json:"Timestamps"`
	// Assets assets to display on the player's profile
	Assets ActivityAssets `json:"Assets"`
	// Party information about the player's party
	Party ActivityParty `json:"Party"`
	// Secrets secret passwords for joining and spectating the player's game
	Secrets ActivitySecrets `json:"Secrets"`
	// Instance whether this activity is an instanced context, like a match
	Instance bool `json:"Instance"`
}

type ActivityTimestamps struct {
	// Start unix timestamp - send this to have an "elapsed" timer
	Start int64 `json:"Start"`
	// End unix timestamp - send this to have a "remaining" timer
	End int64 `json:"End"`
}

type ActivityAssets struct {
	// LargeImage keyname of an asset to display
	LargeImage string `json:"LargeImage"`
	// LargeText hover text for the large image
	LargeText string `json:"LargeText"`
	// SmallImage keyname of an asset to display
	SmallImage string `json:"SmallImage"`
	// SmallText hover text for the small image
	SmallText string `json:"SmallText"`
}

type ActivityParty struct {
	// Id a unique identifier for this party
	Id string `json:"Id"`
	// Size info about the size of the party
	Size PartySize `json:"Size"`
}

type PartySize struct {
	// CurrentSize the current size of the party
	CurrentSize int32 `json:"CurrentSize"`
	// MaxSize the max possible size of the party
	MaxSize int32 `json:"MaxSize"`
}

type ActivitySecrets struct {
	// Match unique hash for the given match context
	Match string `json:"Match"`
	// Join unique hash for chat invites and Ask to Join
	Join string `json:"Join"`
	// Spectate unique hash for Spectate button
	Spectate string `json:"Spectate"`
}