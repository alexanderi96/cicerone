package db

import (
	"errors"
	"log"
	"github.com/alexanderi96/cicerone/types"
)

//CreateUser will create a new user, take as input the parameters and
//insert them into the database
func CreateUser(u types.User) error {
	switch u := u.(type) {
	default:
		log.Println("Undefined user type")
	case types.Cicerone:
		err := gQuery("insert into Utenti(NomeUtente, CognomeUtente, SessoUtente, DataNascitaUtente, EmailUtente, PasswordUtente) values(?,?,?,?,?,?)", u.Nome, u.Cognome, u.Sesso, u.DataNascita, u.Email, u.Password)
		if err != nil {
			return err
		} else {
			u.IdUtente = GetUserID(u.Email)
			if u.IdUtente < 0 {
				return errors.New("Invalid User Id")
			}
			err := gQuery("insert into Ciceroni(IdCicerone, CodiceFiscaleCicerone, TelefonoCicerone, IbanCicerone) values (?,?,?,?)", u.IdUtente, u.Tel, u.Iban, u.CodFis)
			if err != nil {
				return err
			}
		}
	case types.Globetrotter:
		err := gQuery("insert into Utenti(NomeUtente, CognomeUtente, SessoUtente, DataNascitaUtente, EmailUtente, PasswordUtente) values(?,?,?,?,?,?)", u.Nome, u.Cognome, u.Sesso, u.DataNascita, u.Email, u.Password)
		return err
	}
	return nil
}

//ValidUser will check if the user exists in db and if exists if the username password
//combination is valid
func ValidUser(email, password string) bool {
	var passwordFromDB string
	plainSQL := "select PasswordUtente from Utenti where EmailUtente = ?"
	log.Print("validating user ", email)
	rows := database.query(plainSQL, email)

	defer rows.Close()
	log.Println("sus")
	if rows.Next() {
		log.Println("sus")
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
func GetUserID(email string) int {
	var userID int
	userSQL := "select IdUtente from Utenti where EmailUtente=?"
	rows := database.query(userSQL, email)

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&userID)
		if err != nil {
			log.Println(err)
			return -1
		}
	}
	return userID
}

//TODO: currently I'm unable to determin if an user is cicerone or not. Must fix that
func GetUserInfo(email string) (u types.User){
	if IsCicerone(email) {
		user := types.Cicerone{}
		userSQL := "select IdUtente, NomeUtente, CognomeUtente, SessoUtente, DataNascitaUtente, EmailUtente, CodiceFiscaleCicerone, TelefonoCicerone, IbanCicerone from Utenti join Ciceroni on Utenti.IdUtente = Ciceroni.IdCicerone where EmailUtente = ?"
		rows := database.query(userSQL, email)
		defer rows.Close() //must defer after every database interaction

		if rows.Next() {
			err := rows.Scan(&user.IdUtente, &user.Nome, &user.Cognome, &user.Sesso, &user.DataNascita, &user.Email, &user.CodFis, &user.Tel, &user.Iban)	
			if err != nil {
				log.Println(err)
				return types.Cicerone{}
			}
			return user
		}

	} else {
		user := types.Globetrotter{}
		userSQL := "select IdUtente, NomeUtente, CognomeUtente, SessoUtente, DataNascitaUtente, EmailUtente from Utenti where EmailUtente = ?"
		rows := database.query(userSQL, email)
		defer rows.Close() //must defer after every database interaction

		if rows.Next() {
			err := rows.Scan(&user.IdUtente, &user.Nome, &user.Cognome, &user.Sesso, &user.DataNascita, &user.Email)	
			if err != nil {
				log.Println(err)
				return types.Globetrotter{}
			}
		}
		return user
	}
	return
}

func AddCicerone(uid, tel int, iban, fcode string) error {
	err := gQuery("insert into Ciceroni(IdCicerone, CodiceFiscaleCicerone, TelefonoCicerone, IbanCicerone) values (?,?,?,?)", uid, tel, iban, fcode)
	return err
}

func IsCicerone(email string) (false bool) {
	var uidFromDB int
	uid := GetUserID(email)
	if uid == -1 {
		log.Println("Unable to determin user Id")
		return
	}
	userSQL := "select IdCicerone from Ciceroni where IdCicerone=?"
	rows := database.query(userSQL, uid)

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&uidFromDB)
		if err != nil {
			log.Println(err)
			return
		}
		return true
	}
	return
}

func DeleteSelectedUser(email, password string) (e error) {
	if ValidUser(email, password) {
		userSQL := "delete from Utente where email = ?"
		e = gQuery(userSQL, email)
	} else {
		e = errors.New("Invalid User")
	}
	return
}