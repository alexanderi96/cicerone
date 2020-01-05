package views

/*Holds the fetch related view handlers*/

import (
	"html/template"
	"log"
	"net/http"
	"time"
	"strings"
	"github.com/alexanderi96/cicerone/sessions"
)

var templates *template.Template
var homeTemplate *template.Template
var loginTemplate *template.Template
var profileTemplate *template.Template
var eventTemplate *template.Template

var err error

/*
I'm going to deprecate this kinda generic DisplayPage function, in order to restore page specific functions
*/
func DisplayPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		path :=  strings.Split(r.URL.Path, "/")
		usr := sessions.GetCurrentUser(r)
		c, err := loadContext(usr, path)
		if err != nil {
			log.Println("Internal server error retriving context")
			http.Redirect(	w, r, "/", http.StatusInternalServerError)
		} else {
			c.CSRFToken = "abcd" //I need that to utilize cookies. must implement md5 token
			expiration := time.Now().Add(365 * 24 * time.Hour)
			cookie := http.Cookie{Name: "csrftoken", Value: "abcd", Expires: expiration}
			http.SetCookie(w, &cookie)
			
			switch path[1]{
			case "myprofile":
				profileTemplate.Execute(w, c)
			case "event":
				eventTemplate.Execute(w, c)
			default:
				homeTemplate.Execute(w, c)
			}
		}
	}
}

func HomeFunction(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		c.Utente = db.GetUserInfo(usr)

		log.Println("Getting events")
		if c.Events, e = db.GetEvents(); e != nil {
			log.Println("Internal server error retriving context: " + e)
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
