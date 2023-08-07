package entities

type Logs struct {
	Message      string `json:"message,omitempty"`
	StatusCode   string `json:"statusCode,omitempty"`
	Path         string `json:"path,omitempty"`
	Response     string `json:"response,omitempty"`
	QueryParams  string `json:"query_params,omitempty"`
	BodyParams   string `json:"body_params,omitempty"`
	Method       string `json:"method,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
}
