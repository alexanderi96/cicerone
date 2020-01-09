package views

import (
	"log"
	"strconv"
	"net/http"	
	"github.com/alexanderi96/cicerone/db"
	"github.com/alexanderi96/cicerone/sessions"
	"github.com/alexanderi96/cicerone/utils"
	"github.com/alexanderi96/cicerone/types"
)

func SignUpFunc(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Redirect(w, r, "/login/", http.StatusBadRequest)
		return
	}
	r.ParseForm()

	var u types.User

	if r.Form.Get("cicerone") != "on" {
		u = parseGlobe(r)
	} else {
		u = parseCice(r)
	}
	
	e = db.CreateUser(u)
	if e != nil {
		http.Error(w, "Unable to sign user up", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/login/", 302)
	}
}

func GoCicerone(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	r.ParseForm()

	//must implement more security checks
	tel, _ := strconv.Atoi(r.Form.Get("tel"))
	iban := r.Form.Get("iban")
	fcode := r.Form.Get("fcode")

	uid := db.GetUserID(sessions.GetCurrentUser(r))

	e = db.AddCicerone(uid, tel, iban, fcode)
	if e != nil {
		http.Error(w, "Unable to register as Cicerone", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}


func DeleteMyAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}

	r.ParseForm()

	if e := db.DeleteSelectedUser(r.Form.Get("email"), r.Form.Get("password")); e != nil {
		//TODO: handle better this behaviour
		http.Redirect(w, r, "/myprofile/", http.StatusUnauthorized)
	} else {
		log.Println("sas")
		http.Redirect(w, r, "/logout/", 302)
	}
}

func UpdateAccountInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/myprofile/", http.StatusUnauthorized)
		return
	}
	
	uid := db.GetUserID(sessions.GetCurrentUser(r))
	if db.IsCicerone(uid) {
		if e = updateCice(r); e != nil {
			log.Println("(uid: ", uid, ") error updating user information, keeping old details\nError: ", e)
			http.Redirect(w, r, "/myprofile/", http.StatusInternalServerError)
		}
	} else if e = updateGlobe(r); e != nil {
		log.Println("(uid: ", uid, ") error updating user information, keeping old details\nError: ", e)
		http.Redirect(w, r, "/myprofile/", http.StatusInternalServerError)
	}

	//let's make the user log back in in order to load the new info (must find a better way)
	http.Redirect(w, r, "/logout/", 302)	
}

//carico in memoria il vecchio utente in modo da confrontare i campi
//(metodo poco elegante, ma efficace)
func parseCice(r *http.Request) (c types.Cicerone) {
	r.ParseForm()

	tel, _ := strconv.Atoi(r.Form.Get("tel"))
	
	c = types.Cicerone{
		types.Globetrotter{
			0,
			r.Form.Get("name"),
			r.Form.Get("surname"),
			r.Form.Get("gender"),
			utils.DateToUnix(r.Form.Get("bdate")),
			r.Form.Get("email"),

			//TODO: must hash the passwd
			r.Form.Get("password")},
		r.Form.Get("fcode"),
		tel,
		r.Form.Get("iban")}

	return
}

func parseGlobe(r *http.Request) (g types.Globetrotter) {
	r.ParseForm()

	g = types.Globetrotter{
		0,
		r.Form.Get("name"),
		r.Form.Get("surname"),
		r.Form.Get("gender"),
		utils.DateToUnix(r.Form.Get("bdate")),
		r.Form.Get("email"),
		r.Form.Get("password")}

	return
}

func updateCice(r *http.Request) (e error) {
	c := parseCice(r)
	oc, _ := db.GetUserInfo(sessions.GetCurrentUser(r))

	//in order to use the interface without problems wee need to use the "type assertion"
	c.IdUtente = oc.(types.Cicerone).IdUtente
	if c.Nome == "" {
		c.Nome = oc.(types.Cicerone).Nome
	}
	if c.Cognome == "" {
		c.Cognome = oc.(types.Cicerone).Cognome
	}
	if c.Sesso == "" {
		c.Sesso = oc.(types.Cicerone).Sesso
	}
	if c.DataNascita == 0 {
		c.DataNascita = oc.(types.Cicerone).DataNascita
	}
	if c.Email == "" {
		c.Email = oc.(types.Cicerone).Email
	}
	if c.Password == "" {
		c.Password = oc.(types.Cicerone).Password
	}
	if c.CodFis == "" {
		c.CodFis = oc.(types.Cicerone).CodFis
	}
	if c.Tel == 0 {
		c.Tel = oc.(types.Cicerone).Tel
	}
	if c.Iban == "" {
		c.Iban = oc.(types.Cicerone).Iban
	}
	
	//do un id utente negativo alle vecchie info in modo da mantenerle come bk
	//oc.InvertId()

	return SwapInfo(c)
}

func updateGlobe(r *http.Request) (e error){

	g := parseGlobe(r)

	og, _ := db.GetUserInfo(sessions.GetCurrentUser(r))

	
	
	g.IdUtente = og.(types.Globetrotter).IdUtente
	if g.Nome == "" {
		g.Nome = og.(types.Globetrotter).Nome
	}
	if g.Cognome == "" {
		g.Cognome = og.(types.Globetrotter).Cognome
	}
	if g.Sesso == "" {
		g.Sesso = og.(types.Globetrotter).Sesso
	}
	if g.DataNascita == 0 {
		g.DataNascita = og.(types.Globetrotter).DataNascita
	}
	if g.Email == "" {
		g.Email = og.(types.Globetrotter).Email
	}
	if g.Password == "" {
		g.Password = og.(types.Globetrotter).Password
	}
	log.Println("\nBEFORE og: ", og, "\n", "g: ", g)

	return SwapInfo(g)
}

func SwapInfo(u types.User) (e error) {
	//let's first create the backup user with the inverted user id
	if e = db.InvertUserId(u.GetId()); e == nil {
		e = db.CreateUser(u)
	}
	return
}
