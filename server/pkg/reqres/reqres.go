package reqres

type LoginRequest struct {
	UserId string `json:"user_id"`
	UserPw string `json:"user_pw"`
}

type LoginResponse struct {
	ErrorCode int    `json:"error_code"`
	UserId    string `json:"user_id"`
	Token     string `json:"token"`
}

type JoinRequest struct {
	UserId string `json:"user_id"`
	UserPw string `json:"user_pw"`
}

type JoinResponse struct {
	ErrorCode int `json:"error_code"`
}

type MiningRequest struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

// 캔 코인 수
type MiningResponse struct {
	ErrorCode int `json:"error_code"`
	Coin      int `json:"coin"`
}

type UserRequest struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type UserResponse struct {
	ErrorCode int `json:"error_code"`
	Coin      int `json:"coin"`
	MaxCoin   int `json:"max_coin"`
}
