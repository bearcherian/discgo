package structs

type Team struct {
	// Icon a hash of the image of the team's icon
	Icon string `json:"icon"`
	// Id the unique id of the team
	Id string `json:"id"`
	// Members the members of the team
	Members []TeamMember `json:"members"`
	// Name the name of the team
	Name string `json:"name"`
	// OwnerUserId the user id of the current team owner
	OwnerUserId string `json:"owner_user_id"`
}

type TeamMembers struct {
	// MembershipState the user's [membership state](#DOCS_TOPICS_TEAMS/data-models-membership-state-enum) on the team
	MembershipState int `json:"membership_state"`
	// Permissions will always be `["*"]`
	Permissions []string `json:"permissions"`
	// TeamId the id of the parent team of which they are a member
	TeamId string `json:"team_id"`
	// User the avatar, discriminator, id, and username of the user
	User User `json:"user"`
}