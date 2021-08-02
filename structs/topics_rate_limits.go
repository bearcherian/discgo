package structs

type RateLimitResponse struct {
	// Message A message saying you are being rate limited.
	Message string `json:"message"`
	// RetryAfter The number of seconds to wait before submitting another request.
	RetryAfter float32 `json:"retry_after"`
	// Global A value indicating if you are being globally rate limited or not
	Global bool `json:"global"`
}