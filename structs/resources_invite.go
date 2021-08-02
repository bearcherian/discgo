package structs

type Invite struct {
	// Code the invite code (unique ID)
	Code string `json:"code"`
	// Guild the guild this invite is for
	Guild Guild `json:"guild"`
	// Channel the channel this invite is for
	Channel Channel `json:"channel"`
	// Inviter the user who created the invite
	Inviter User `json:"inviter"`
	// TargetType the [type of target](#DOCS_RESOURCES_INVITE/invite-object-invite-target-types) for this voice channel invite
	TargetType int `json:"target_type"`
	// TargetUser the user whose stream to display for this voice channel stream invite
	TargetUser User `json:"target_user"`
	// TargetApplication the embedded application to open for this voice channel embedded application invite
	TargetApplication Application `json:"target_application"`
	// ApproximatePresenceCount approximate count of online members, returned from the `GET /invites/<code>` endpoint when `with_counts` is `true`
	ApproximatePresenceCount int `json:"approximate_presence_count"`
	// ApproximateMemberCount approximate count of total members, returned from the `GET /invites/<code>` endpoint when `with_counts` is `true`
	ApproximateMemberCount int `json:"approximate_member_count"`
	// ExpiresAt the expiration date of this invite, returned from the `GET /invites/<code>` endpoint when `with_expiration` is `true`
	ExpiresAt time.Time `json:"expires_at"`
	// StageInstance stage instance data if there is a [public Stage instance](#DOCS_RESOURCES_STAGE_INSTANCE) in the Stage channel this invite is for
	StageInstance InviteStageInstance `json:"stage_instance"`
}

type InviteMetadata struct {
	// Uses number of times this invite has been used
	Uses int `json:"uses"`
	// MaxUses max number of times this invite can be used
	MaxUses int `json:"max_uses"`
	// MaxAge duration (in seconds) after which the invite expires
	MaxAge int `json:"max_age"`
	// Temporary whether this invite only grants temporary membership
	Temporary bool `json:"temporary"`
	// CreatedAt when this invite was created
	CreatedAt time.Time `json:"created_at"`
}

type InviteStageInstance struct {
	// Members the members speaking in the Stage
	Members []GuildMember `json:"members"`
	// ParticipantCount the number of users in the Stage
	ParticipantCount int `json:"participant_count"`
	// SpeakerCount the number of users speaking in the Stage
	SpeakerCount int `json:"speaker_count"`
	// Topic the topic of the Stage instance (1-120 characters)
	Topic string `json:"topic"`
}