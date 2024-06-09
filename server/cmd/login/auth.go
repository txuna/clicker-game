package main

import (
	"clicker/internal/database"
	"clicker/pkg/reqres"

	"github.com/google/uuid"
)

func (ls *LoginServer) DoJoin(req *reqres.JoinRequest) int {

	hasAccount, err := database.HasAccount(ls.MysqlConn, req.UserId)
	if err != nil {
		return reqres.ERROR_INVALID_REQUEST
	}

	if hasAccount {
		return reqres.ERROR_ALREADY_EXIST_USER
	}

	idx, err := database.InsertAccount(ls.MysqlConn, req.UserId, req.UserPw)
	if err != nil {
		return reqres.ERROR_INTERNAL_SERVER
	}

	// player data 생성
	_ = idx

	return reqres.ERROR_NONE
}

func (ls *LoginServer) DoLogin(req *reqres.LoginRequest) (string, int) {

	account, err := database.FindAccount(ls.MysqlConn, req.UserId)
	if err != nil {
		return "", reqres.ERROR_INTERNAL_SERVER
	}

	result := database.ComparePassword(req.UserPw, account.UserPw)
	if !result {
		return "", reqres.ERROR_INVALID_USER
	}

	// 토큰 발급
	token, err := uuid.NewV7()
	if err != nil {
		return "", reqres.ERROR_FAILED_GENERATE_TOKEN
	}

	// 레디스에 로그인 정보 저장
	return token.String(), reqres.ERROR_NONE
}
