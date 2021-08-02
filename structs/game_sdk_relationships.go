package structs

type Relationship struct {
	// Type what kind of relationship it is
	Type RelationshipType `json:"Type"`
	// User the user the relationship is for
	User User `json:"User"`
	// Presence that user's current presence
	Presence Presence `json:"Presence"`
}

type Presence struct {
	// Status the user's current online status
	Status Status `json:"Status"`
	// Activity the user's current activity
	Activity Activity `json:"Activity"`
}