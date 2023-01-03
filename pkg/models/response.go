package models

const (
	StatusOk    = "ok"
	StatusError = "error"
)

type NetResponse struct {
	H          map[string]any
	StatusCode int
}

func (net NetResponse) Build() NetResponse {
	net.H = make(map[string]any)
	return net
}

func (net NetResponse) SetStatus(statusCode int, status string, message string) {
	net.StatusCode = statusCode
	net.H["status"] = status
	net.H["message"] = message
}

func (net NetResponse) Set(key string, value any) {
	net.H[key] = value
}

func (net NetResponse) Generate() (int, map[string]any) {
	return net.StatusCode, net.H
}
