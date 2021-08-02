package structs

type Lobby struct {
	// Id the unique id of the lobby
	Id int64 `json:"Id"`
	// Type if the lobby is public or private
	Type LobbyType `json:"Type"`
	// OwnerId the userId of the lobby owner
	OwnerId int64 `json:"OwnerId"`
	// Secret the password to the lobby
	Secret string `json:"Secret"`
	// Capacity the max capacity of the lobby
	Capacity uint32 `json:"Capacity"`
	// Locked whether or not the lobby can be joined
	Locked bool `json:"Locked"`
}

type SearchFilter struct {
	// Key the metadata key to search
	Key string `json:"key"`
	// Value the value of the metadata key to validate against
	Value string `json:"value"`
	// Cast the type to cast `value` as
	Cast SearchCast `json:"cast"`
	// Comparison how to compare the metadata values
	Comparison SearchComparison `json:"comparison"`
}

type SearchSort struct {
	// Key the metadata key on which to sort lobbies that meet the search criteria
	Key string `json:"key"`
	// Cast the type to cast `value` as
	Cast SearchCast `json:"cast"`
	// NearValue the value around which to sort the key
	NearValue string `json:"near_value"`
}