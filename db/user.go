package db

import (
	"log"
	"github.com/alexanderi96/cicerone/types"
)

//CreateUser will create a new user, take as input the parameters and
//insert it into database
func CreateUser(name, surname, username, password, email string) error {
	err := gQuery("insert into users(name, surname, username, passwd_h, email, cicerone) values(?,?,?,?,?,?)", name, surname, username, password, email, false)
	return err
}

//ValidUser will check if the user exists in db and if exists if the username password
//combination is valid
func ValidUser(username, password string) bool {
	var passwordFromDB string
	plainSQL := "select passwd_h from users where username=?"
	log.Print("validating user ", username)
	rows := database.query(plainSQL, username)

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&passwordFromDB)
		if err != nil {
			return false
		}
	}
	//If the password matches, return true
	if password == passwordFromDB {
		return true
	}
	//by default return false
	return false
}

//GetUserID will get the user's ID from the database
func GetUserID(username string) (int, error) {
	var userID int
	userSQL := "select uid from users where username=?"
	rows := database.query(userSQL, username)

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&userID)
		if err != nil {
			return -1, err
		}
	}
	return userID, nil
}

func GetUserInfo(username string) (types.User, error) {
	var user types.User

	userSQL := "select uid, name, surname, username, email, passwd_h, cicerone from users where username=?"
	rows := database.query(userSQL, username)

	defer rows.Close() //must defer after every database interaction
	if rows.Next() {

		err := rows.Scan(&user.Uid, &user.Name, &user.Surname, &user.Username, &user.Email, &user.Passwd_h, &user.Cicerone)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}
