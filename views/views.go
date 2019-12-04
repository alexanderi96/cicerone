package views

/*Holds the fetch related view handlers*/

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"gitlab.com/alexanderi96/cicerone/db"
	"gitlab.com/alexanderi96/cicerone/sessions"
	"gitlab.com/alexanderi96/cicerone/types"

)

var templates *template.Template
var homeTemplate *template.Template
var loginTemplate *template.Template
var profileTemplate *template.Template

var err error

//HomePageFunc is used to handle the "/" URL which is the default page
//TODO add http404 error
func HomePageFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {		
		c, err := loadContext(r)
		if err != nil {
			log.Println("Internal server error retriving context")
			http.Redirect(	w, r, "/", http.StatusInternalServerError)
		} else {
			c.CSRFToken = "abcd" //I need that to utilize cookies. must implement md5 token
			expiration := time.Now().Add(365 * 24 * time.Hour)
			cookie := http.Cookie{Name: "csrftoken", Value: "abcd", Expires: expiration}
			http.SetCookie(w, &cookie)
			homeTemplate.Execute(w, c)
		}
	}
}

func MyProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		c, err := loadContext(r)
		if err != nil {
			log.Println("Internal server error retriving context")
			http.Redirect(	w, r, "/", http.StatusInternalServerError)
		} else {
			c.CSRFToken = "abcd" //I need that to utilize cookies. must implement md5 token
			expiration := time.Now().Add(365 * 24 * time.Hour)
			cookie := http.Cookie{Name: "csrftoken", Value: "abcd", Expires: expiration}
			http.SetCookie(w, &cookie)
			profileTemplate.Execute(w, c)
		}
	}
}

func loadContext(r *http.Request) (*types.Context, error) {
	var c types.Context

	c.Utente, err = db.GetUserInfo(sessions.GetCurrentUser(r))

	if err != nil {
		log.Println("Error retriving user info: ", err)
		return nil, err
	} else {
		c.IsCicerone, err = db.IsCicerone(c.Utente.IdUtente)
		if err != nil {
			log.Println("Error establishing IfCicerone")
			return nil, err
		} else {
			c.Events, err = db.GetEvents()
			return &c, nil
		}
	}
	
}  