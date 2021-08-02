package structs

type Device struct {
	// Type the type of device
	Type [device type](#DOCS_TOPICS_CERTIFIED_DEVICES/models-device-type) `json:"type"`
	// Id the device's Windows UUID
	Id string `json:"id"`
	// Vendor the hardware vendor
	Vendor Vendor `json:"vendor"`
	// Model the model of the product
	Model Model `json:"model"`
	// Related UUIDs of related devices
	Related []string `json:"related"`
	// EchoCancellation if the device's native echo cancellation is enabled
	EchoCancellation bool `json:"echo_cancellation"`
	// NoiseSuppression if the device's native noise suppression is enabled
	NoiseSuppression bool `json:"noise_suppression"`
	// AutomaticGainControl if the device's native automatic gain control is enabled
	AutomaticGainControl bool `json:"automatic_gain_control"`
	// HardwareMute if the device is hardware muted
	HardwareMute bool `json:"hardware_mute"`
}

type Vendor struct {
	// Name name of the vendor
	Name string `json:"name"`
	// Url url for the vendor
	Url string `json:"url"`
}

type Model struct {
	// Name name of the model
	Name string `json:"name"`
	// Url url for the model
	Url string `json:"url"`
}