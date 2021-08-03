package structs

type AuditLogEventType int

const (
	AuditLogEventGuildUpdate            AuditLogEventType = 1
	AuditLogEventChannelCreate          AuditLogEventType = 10
	AuditLogEventChannelUpdate          AuditLogEventType = 11
	AuditLogEventChannelDelete          AuditLogEventType = 12
	AuditLogEventChannelOverwriteCreate AuditLogEventType = 13
	AuditLogEventChannelOverwriteUpdate AuditLogEventType = 14
	AuditLogEventChannelOverwriteDelete AuditLogEventType = 15
	AuditLogEventMemberKick             AuditLogEventType = 20
	AuditLogEventMemberPrune            AuditLogEventType = 21
	AuditLogEventMemberBanAdd           AuditLogEventType = 22
	AuditLogEventMemberBanRemove        AuditLogEventType = 23
	AuditLogEventMemberUpdate           AuditLogEventType = 24
	AuditLogEventMemberRoleUpdate       AuditLogEventType = 25
	AuditLogEventMemberMove             AuditLogEventType = 26
	AuditLogEventMemberDisconnect       AuditLogEventType = 27
	AuditLogEventBotAdd                 AuditLogEventType = 28
	AuditLogEventRoleCreate             AuditLogEventType = 30
	AuditLogEventRoleUpdate             AuditLogEventType = 31
	AuditLogEventRoleDelete             AuditLogEventType = 32
	AuditLogEventInviteCreate           AuditLogEventType = 40
	AuditLogEventInviteUpdate           AuditLogEventType = 41
	AuditLogEventInviteDelete           AuditLogEventType = 42
	AuditLogEventWebhookCreate          AuditLogEventType = 50
	AuditLogEventWebhoolUpdate          AuditLogEventType = 51
	AuditLogEventWebhookDelete          AuditLogEventType = 52
	AuditLogEventEmojiCreate            AuditLogEventType = 60
	AuditLogEventEmojiUpdate            AuditLogEventType = 61
	AuditLogEventEmojiDelete            AuditLogEventType = 62
	AuditLogEventMessageDelete          AuditLogEventType = 72
	AuditLogEventMessageBulkDelete      AuditLogEventType = 73
	AuditLogEventMessagePin             AuditLogEventType = 74
	AuditLogEventMessageUnpin           AuditLogEventType = 75
	AuditLogEventIntegrationCreate      AuditLogEventType = 80
	AuditLogEventIntegrationUpdate      AuditLogEventType = 81
	AuditLogEventIntegrationDelete      AuditLogEventType = 82
	AuditLogEventStageInstaceCreate     AuditLogEventType = 83
	AuditLogEventStageInstanceUpdate    AuditLogEventType = 84
	AuditLogEventStageInstanceDelete    AuditLogEventType = 85
	AuditLogEventStickerCreate          AuditLogEventType = 90
	AuditLogEventStickerUpdate          AuditLogEventType = 91
	AuditLogEventStickerDelete          AuditLogEventType = 92
)

type AuditLog struct {
	// Webhooks list of webhooks found in the audit log
	Webhooks []Webhook `json:"webhooks"`
	// Users list of users found in the audit log
	Users []User `json:"users"`
	// AuditLogEntries list of audit log entries
	AuditLogEntries []AuditLogEntry `json:"audit_log_entries"`
	// Integrations list of partial integration objects
	Integrations []Integration `json:"integrations"`
	// Threads list of threads found in the audit log\*
	Threads []Channel `json:"threads"`
}

type AuditLogEntry struct {
	// TargetId id of the affected entity (webhook, user, role, etc.)
	TargetId string `json:"target_id"`
	// Changes changes made to the target_id
	Changes []AuditLogChange `json:"changes"`
	// UserId the user who made the changes
	UserId string `json:"user_id"`
	// Id id of the entry
	Id string `json:"id"`
	// ActionType type of action that occurred
	ActionType AuditLogEventType `json:"action_type"`
	// Options additional info for certain action types
	Options AuditEntryInfo `json:"options"`
	// Reason the reason for the change (0-512 characters)
	Reason string `json:"reason"`
}

type AuditEntryInfo struct {
	DeleteMemberDays string `json:"delete_member_days"`
	MembersRemoved   string `json:"members_removed"`
	ChannelID        string `json:"channel_id"`
	MessageID        string `json:"message_id"`
	Count            string `json:"count"`
	Id               string `json:"id"`
	Type             string `json:"type"`
	RoleName         string `json:"role_name"`
}

type AuditLogChange struct {
	// NewValue new value of the key
	NewValue interface{} `json:"new_value"`
	// OldValue old value of the key
	OldValue interface{} `json:"old_value"`
	// Key name of audit log [change key](#DOCS_RESOURCES_AUDIT_LOG/audit-log-change-object-audit-log-change-key)
	Key string `json:"key"`
}
