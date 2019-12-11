package db

import (
	"log"
	"github.com/alexanderi96/cicerone/types"
)

//CreateUser will create a new user, take as input the parameters and
//insert it into database
func CreateUser(u types.Utente) error {
	err := gQuery("insert into Utenti(NomeUtente, CognomeUtente, SessoUtente, DataNascitaUtente, EmailUtente, PasswordUtente) values(?,?,?,?,?,?)", u.Nome, u.Cognome, u.Sesso, u.DataNascita, u.Email, u.Password)
	return err
}

//ValidUser will check if the user exists in db and if exists if the username password
//combination is valid
func ValidUser(email, password string) bool {
	var passwordFromDB string
	plainSQL := "select PasswordUtente from Utenti where EmailUtente=?"
	log.Print("validating user ", email)
	rows := database.query(plainSQL, email)

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&passwordFromDB)
		if err != nil {
			return false
		}
		//If the password matches, return true
		if password == passwordFromDB {
			return true
		}
	}
	
	//by default return false
	return false
}

//GetUserID will get the user's ID from the database
func GetUserID(email string) (int, error) {
	var userID int
	userSQL := "select IdUtente from Utenti where EmailUtente=?"
	rows := database.query(userSQL, email)

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&userID)
		if err != nil {
			return -1, err
		}
	}
	return userID, nil
}

//TODO: currently I'm unable to determin if an user is cicerone or not. Must fix that
func GetUserInfo(email string, usr chan types.Utente) {
	var user types.Utente 
	userSQL := "select IdUtente, NomeUtente, CognomeUtente, SessoUtente, DataNascitaUtente, EmailUtente from Utenti where EmailUtente=?"
	rows := database.query(userSQL, email)

	defer rows.Close() //must defer after every database interaction
	if rows.Next() {
		_ = rows.Scan(&user.IdUtente, &user.Nome, &user.Cognome, &user.Sesso, &user.DataNascita, &user.Email)
		
		userSQL = "select IdCicerone, TelefonoCicerone, CodiceFiscaleCicerone, IbanCicerone from Ciceroni where IdCicerone=?"
		rows = database.query(userSQL, user.IdUtente)
		defer rows.Close()
		if rows.Next() {
			_ = rows.Scan(&user.Cicerone.IdCicerone, &user.Cicerone.Tel, &user.Cicerone.CodFis, &user.Cicerone.Iban)	
		}
	}
	log.Println(user)
	usr <- user
}

func AddCicerone(uid, tel int, iban, fcode string) error {
	err := gQuery("insert into Ciceroni(IdCicerone, CodiceFiscaleCicerone, TelefonoCicerone, IbanCicerone) values (?,?,?,?)", uid, tel, iban, fcode)
	return err
}

func IsCicerone(email string) (bool, error) {
	var uidFromDB int
	uid, err := GetUserID(email)
	if err != nil {
		return false, err
	}
	userSQL := "select IdCicerone from Ciceroni where IdCicerone=?"
	rows := database.query(userSQL, uid)

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&uidFromDB)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}