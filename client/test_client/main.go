package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
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

func ranking(userid, token string) RankingResponse {
	req := RankingRequest{
		UserId: userid,
		Token:  token,
	}

	resp := request(&req, GAME_URL+"ranking")
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := RankingResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			run(fmt.Sprintf("%d_user", i), fmt.Sprintf("%d_password", i))
		}(i)
		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()

	fmt.Println(calculateStats(durations))
}

var durations []time.Duration
var mutex sync.Mutex

func run(userid, password string) {
	log.Printf("Join: %v\n", join(userid, password))
	v := login(userid, password)
	log.Printf("login: %v\n", v.ErrorCode)

	for i := 0; i < 30; i++ {
		start := time.Now()
		res := mining(userid, v.Token)
		log.Printf("mining: %d coin - userid: %s\n", res.Coin, userid)
		elapsed := time.Since(start)
		store(elapsed)
		//fmt.Printf("Execution time: %s\n", elapsed)
	}
}

func store(e time.Duration) {
	mutex.Lock()
	defer mutex.Unlock()
	durations = append(durations, e)
}

// Calculate average, maximum, and minimum durations
func calculateStats(durations []time.Duration) (average, max, min time.Duration) {
	if len(durations) == 0 {
		return
	}

	var total time.Duration
	max = durations[0]
	min = durations[0]

	for _, duration := range durations {
		total += duration
		if duration > max {
			max = duration
		}
		if duration < min {
			min = duration
		}
	}

	average = total / time.Duration(len(durations))
	return
}
