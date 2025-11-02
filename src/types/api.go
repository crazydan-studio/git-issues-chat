package types

// Response represents the response data structure
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Msg     string      `json:"msg,omitempty"`
}