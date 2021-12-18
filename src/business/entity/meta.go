package entity


type ErrorMessage struct {
	Code     string `json:"code"`
	Message string `json:"message"`
}

type Meta struct {
	Path       string `json:"path"`
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
	Error      ErrorMessage `json:"error,omitempty"`
	Timestamp  string `json:"timestamp"`
}

type HTTPErrResp struct {
	Meta Meta `json:"metadata"`
}
type HTTPEmptyResp struct {
	Meta Meta `json:"metadata"`
}
