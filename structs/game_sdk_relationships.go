package structs

type RelationshipType string
type Status int

const (
	RelationshipNone            = "None"
	RelationshipFriend          = "Friend"
	RelationshipBlocked         = "Blocked"
	RelationshipPendingIncoming = "PendingIncoming"
	RelationshipPendingOutgoing = "PendingOutgoing"
	RelationshipImplicit        = "Implicit"

	StatusOffline      = 0
	StatusOnline       = 1
	StatusIdle         = 2
	StatusDoNotDisturb = 3
)

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
