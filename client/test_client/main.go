package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	LOGIN_URL = "http://127.0.0.1:9001/"
	GAME_URL  = "http://127.0.0.1:9003/"
)

var TOKEN = ""
var USERID = ""

func request(v any, url string) *http.Response {
	data, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	buff := bytes.NewBuffer(data)
	resp, err := http.Post(url, "application/json", buff)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func join(userid, userpw string) JoinResponse {
	req := JoinRequest{
		UserId: userid,
		UserPw: userpw,
	}

	resp := request(&req, LOGIN_URL+"join")
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := JoinResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func login(userid, userpw string) LoginResponse {
	req := LoginRequest{
		UserId: userid,
		UserPw: userpw,
	}

	resp := request(&req, LOGIN_URL+"login")
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := LoginResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal(err)
	}

	if res.ErrorCode == 1000 {
		TOKEN = res.Token
		USERID = req.UserId
	}

	return res
}

func user(userid, token string) PlayerResponse {
	req := PlayerRequest{
		UserId: userid,
		Token:  token,
	}

	resp := request(&req, GAME_URL+"user")
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := PlayerResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func mining(userid, token string) MiningResponse {
	req := MiningRequest{
		UserId: userid,
		Token:  token,
	}

	resp := request(&req, GAME_URL+"mining")
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := MiningResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func main() {
	log.Printf("Join: %v\n", join("lemon", "1234"))
	log.Printf("login: %v\n", login("lemon", "1234"))
	log.Printf("user: %v\n", user(USERID, TOKEN))
	for i := 0; i < 30; i++ {
		log.Printf("mining: %v\n", mining(USERID, TOKEN))
		time.Sleep(100 * time.Millisecond)
	}
	log.Printf("user: %v\n", user(USERID, TOKEN))
}
