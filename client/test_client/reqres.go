package main

type JoinRequest struct {
	UserId string `json:"user_id"`
	UserPw string `json:"user_pw"`
}

type JoinResponse struct {
	ErrorCode int `json:"error_code"`
}

type LoginRequest struct {
	UserId string `json:"user_id"`
	UserPw string `json:"user_pw"`
}

type LoginResponse struct {
	ErrorCode int    `json:"error_code"`
	Token     string `json:"token"`
}

type PlayerRequest struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type PlayerResponse struct {
	ErrorCode int `json:"error_code"`
	Coin      int `json:"coin"`
	MaxCoin   int `json:"max_coin"`
}

type MiningRequest struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type MiningResponse struct {
	ErrorCode int `json:"error_code"`
	Coin      int `json:"coin"`
}

type RankingRequest struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type Player struct {
	PlayerId int `json:"player_id"`
	Coin     int `json:"coin"`
	MaxCoin  int `json:"max_coin"`
}

type RankingResponse struct {
	ErrorCode int      `json:"error_code"`
	Players   []Player `json:"players"`
}
