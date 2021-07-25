package transport

type PongRequest struct {
	Data string `json:"data"`
}

type PongResponse struct {
	Data string `json:"pong"`
}
