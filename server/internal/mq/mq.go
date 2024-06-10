package mq

type NatsLoginSearchRequest struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type NatsLoginSearchResponse struct {
	Result bool `json:"result"`
}

type NatsLoginStoreRequest struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}
