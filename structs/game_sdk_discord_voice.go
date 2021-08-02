package structs

type InputModeType int

const (
	InputModeVoiceActivity InputModeType = 0
	InputModePushToTalk    InputModeType = 1
)

type InputMode struct {
	// Type set either VAD or PTT as the voice input mode
	Type InputModeType `json:"Type"`
	// Shortcut the PTT hotkey for the user
	Shortcut string `json:"Shortcut"`
}
