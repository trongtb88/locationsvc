package common

const (
	StatusActive   = 1
	StatusInActive = 0

	StatusActiveStr   = "Active"
	StatusInActiveStr = "InActive"

	HttpHeaderContentType = "Content-Type"
	HttpContentJSON = "application/json"

	DefaultLimit = 10
	DefaultPage = 1
)

func FillStatusInStr(status int) string {
	if status == StatusActive {
		return StatusActiveStr
	}
	return StatusInActiveStr
}

func FillStatus(statusInStr string) int {
	if statusInStr == StatusActiveStr {
		return StatusActive
	}
	return StatusInActive
}

