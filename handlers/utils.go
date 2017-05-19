package handlers

// HTTPResp the typical struct used to send back json responses
type HTTPResp struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Body        string `json:"body"`
}
