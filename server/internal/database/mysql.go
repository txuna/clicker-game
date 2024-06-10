package database

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Id     int64
	UserId string
	UserPw string
}

func DeleteAccount(db *sql.DB, userid string) error {
	query := "DELETE FROM `accounts` where userid = ?;"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userid)
	return err
}

func DeletePlayer(db *sql.DB, playerId int) error {
	query := "DELETE FROM `players` where player_id = ?;"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(playerId)
	return err
}

func InsertPlayer(db *sql.DB, playerId, coin int) error {
	query := "INSERT INTO `players` (player_id, coin) VALUES (?, ?);"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(playerId, coin)
	return err
}

func InsertAccount(db *sql.DB, userId, userPw string) (int64, error) {
	hashPw, err := hashAndSalt([]byte(userPw))
	if err != nil {
		return -1, err
	}

	query := "INSERT INTO `accounts` (userid, userpw) VALUES (?, ?);"
	stmt, err := db.Prepare(query)
	if err != nil {
		return -1, err
	}

	result, err := stmt.Exec(userId, hashPw)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func FindAccount(db *sql.DB, userId string) (Account, error) {
	var account Account
	stmt, err := db.Prepare("SELECT * FROM accounts where userid = ?")
	if err != nil {
		return Account{}, err
	}

	err = stmt.QueryRow(userId).Scan(&account.Id, &account.UserId, &account.UserPw)
	if err != nil {
		return Account{}, err
	}

	return account, err
}

func HasAccount(db *sql.DB, userId string) (bool, error) {
	stmt, err := db.Prepare("SELECT count(*) FROM accounts where userid = ?")
	if err != nil {
		return false, err
	}

	var has bool
	err = stmt.QueryRow(userId).Scan(&has)
	if err != nil {
		return false, err
	}

	return has, nil
}

func ComparePassword(reqPw, hashPw string) bool {
	return comparePasswords(hashPw, []byte(reqPw))
}

func hashAndSalt(pwd []byte) (string, error) {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	return err == nil
}
