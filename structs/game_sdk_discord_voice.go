package structs

type InputMode struct {
	// Type set either VAD or PTT as the voice input mode
	Type InputModeType `json:"Type"`
	// Shortcut the PTT hotkey for the user
	Shortcut string `json:"Shortcut"`
}