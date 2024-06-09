package reqres

type LoginRequest struct {
	UserId string `json:"user_id"`
	UserPw string `json:"user_pw"`
}

type LoginResponse struct {
	ErrorCode int    `json:"error_code"`
	Token     string `json:"token"`
}

type JoinRequest struct {
	UserId string `json:"user_id"`
	UserPw string `json:"user_pw"`
}

type JoinResponse struct {
	ErrorCode int `json:"error_code"`
}
