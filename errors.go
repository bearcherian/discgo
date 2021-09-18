package discgo

const (
	ErrUnableToCreateRequest   = "unable to create HTTP request: %s"
	ErrUnableToCompleteRequest = "unable to complete HTTP request: %s"
	ErrUnableToParseResponse   = "unable to read response body: %s"
)

type ErrorResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Errors  map[string]interface{} `json:"errors"`
}
