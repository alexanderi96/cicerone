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
	
	err := db.CreateUser(u)
	if err != nil {
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

	err := db.AddCicerone(uid, tel, iban, fcode)
	if err != nil {
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
		log.Println(e)
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

	r.ParseForm()

	var u types.User
	uid, _ := strconv.Atoi(r.Form.Get("userid"))

	//TODO: make something to check which user details has been changed maybe using a range function

	if db.IsCicerone(uid) {
		u = parseCice(r)
	} else {
		u = parseGlobe(r)
	}
	log.Println(u)
	if e := db.UpdateUserInfo(uid, u); e != nil {
		log.Println("error updating user information, keeping old details")
		http.Redirect(w, r, "/myprofile/", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/myprofile/", 302)	
	}

	
}

func parseCice(r *http.Request) (u types.Cicerone) {
	tel, _ := strconv.Atoi(r.Form.Get("tel"))

	u = types.Cicerone{
		types.Globetrotter{
			-1,
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

func parseGlobe(r *http.Request) (u types.Globetrotter){
	u = types.Globetrotter{
		-1,
		r.Form.Get("name"),
		r.Form.Get("surname"),
		r.Form.Get("gender"),
		utils.DateToUnix(r.Form.Get("bdate")),
		r.Form.Get("email"),
		r.Form.Get("password")}
	return
}