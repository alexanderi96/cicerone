package db

import (
	"log"
	"github.com/alexanderi96/cicerone/types"
)

//CreateUser will create a new user, take as input the parameters and
//insert it into database
func CreateUser(u types.Utente) error {
	err := gQuery("insert into Utente(Nome, Cognome, Sesso, DataNascita, Email, Password) values(?,?,?,?,?,?)", u.Nome, u.Cognome, u.Sesso, u.DataNascita, u.Email, u.Password)
	return err
}

//ValidUser will check if the user exists in db and if exists if the username password
//combination is valid
func ValidUser(email, password string) bool {
	var passwordFromDB string
	plainSQL := "select Password from Utente where Email=?"
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
	userSQL := "select IdUtente from Utente where Email=?"
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

func GetUserInfo(email string) (*types.Utente, error) {
	var user types.Utente
	
	userSQL := "select IdUtente, Nome, Cognome, Sesso, DataNascita, Email from Utente where Email=?"
	rows := database.query(userSQL, email)

	defer rows.Close() //must defer after every database interaction
	if rows.Next() {

		err := rows.Scan(&user.IdUtente, &user.Nome, &user.Cognome, &user.Sesso, &user.DataNascita, &user.Email)
		if err != nil {
			return &user, err
		}
		user.Cicerone, _ = GetCiceInfo(user.IdUtente)

	}
	return &user, nil
}

func GetCiceInfo(uid int) (*types.Cicerone, error) {
	var cice types.Cicerone

	userSQL := "select IdCicerone, Telefono, Fcode, Iban from Cicerone where IdCicerone=?"
	rows := database.query(userSQL, uid)
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&cice.IdCicerone, &cice.Tel, &cice.CodFis, &cice.Iban)
		if err != nil {
			return &cice, err
		}
	}
	return &cice, nil
}

func AddCicerone(uid, tel int, iban, fcode string) error {
	err := gQuery("insert into Cicerone(IdCicerone, Fcode, Telefono, Iban) values (?,?,?,?)", uid, tel, iban, fcode)
	return err
}

func IsCicerone(uid int) (bool, error) {
	cicerone := false
	var uidFromDB int
	userSQL := "select IdCicerone from Cicerone where IdCicerone=?"
	rows := database.query(userSQL, uid)

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&uidFromDB)
		if err != nil {
			return cicerone, err
		}
		cicerone = true
	}

	return cicerone, nil
}
