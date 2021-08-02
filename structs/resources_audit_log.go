package structs

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
	ActionType [audit log event](#DOCS_RESOURCES_AUDIT_LOG/audit-log-entry-object-audit-log-events) `json:"action_type"`
	// Options additional info for certain action types
	Options [optional audit entry info](#DOCS_RESOURCES_AUDIT_LOG/audit-log-entry-object-optional-audit-entry-info) `json:"options"`
	// Reason the reason for the change (0-512 characters)
	Reason string `json:"reason"`
}

type AuditLogChange struct {
	// NewValue new value of the key
	NewValue [mixed](#DOCS_RESOURCES_AUDIT_LOG/audit-log-change-object-audit-log-change-key) `json:"new_value"`
	// OldValue old value of the key
	OldValue [mixed](#DOCS_RESOURCES_AUDIT_LOG/audit-log-change-object-audit-log-change-key) `json:"old_value"`
	// Key name of audit log [change key](#DOCS_RESOURCES_AUDIT_LOG/audit-log-change-object-audit-log-change-key)
	Key string `json:"key"`
}