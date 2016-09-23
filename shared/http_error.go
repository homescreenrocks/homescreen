package shared

type HttpError struct {
	Message string `json:"message"`
	Error   error  `json:"error,omitempty"`
}
