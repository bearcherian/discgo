package structs

type Achievement struct {
	// ApplicationId the unique id of the application
	ApplicationId int64 `json:"application_id"`
	// Name the name of the achievement as an [achievement locale object](#DOCS_GAME_SDK_ACHIEVEMENTS/achievement-locale-object)
	Name AchievementLocale `json:"name"`
	// Description the user-facing achievement description as an [achievement locale object](#DOCS_GAME_SDK_ACHIEVEMENTS/achievement-locale-object)
	Description AchievementLocale `json:"description"`
	// Secret if the achievement is secret
	Secret bool `json:"secret"`
	// Secure if the achievement is secure
	Secure bool `json:"secure"`
	// Id the unique id of the achievement
	Id int64 `json:"id"`
	// IconHash [the hash of the icon](#DOCS_REFERENCE/image-formatting)
	IconHash string `json:"icon_hash"`
}

type AchievementLocale struct {
	// Default the default locale for the achievement
	Default string `json:"default"`
	// Localizations object of accepted locales as the key and achievement translations as the value
	Localizations map[string]string `json:"localizations"`
}

type UserAchievement struct {
	// UserId the unique id of the user working on the achievement
	UserId int64 `json:"UserId"`
	// AchievementId the unique id of the achievement
	AchievementId int64 `json:"AchievementId"`
	// PercentComplete how far along the user is to completing the achievement (0-100)
	PercentComplete uint8 `json:"PercentComplete"`
	// UnlockedAt the timestamp at which the user completed the achievement (PercentComplete was set to 100)
	UnlockedAt string `json:"UnlockedAt"`
}
