package mq

type NatsLoginSearchResponse struct {
	ErrorCode int
}

type NatsLoginStoreRequest struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}
