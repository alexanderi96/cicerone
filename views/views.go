package views

/*Holds the fetch related view handlers*/

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/alexanderi96/cicerone/db"
	"github.com/alexanderi96/cicerone/sessions"
	"github.com/alexanderi96/cicerone/types"

)

var templates *template.Template
var homeTemplate *template.Template
var loginTemplate *template.Template

var err error

//HomePageFunc is used to handle the "/" URL which is the default page
//TODO add http404 error
func HomePageFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var c types.Context
		c.Events, err = db.GetEvents()
		c.User, err = db.GetUserInfo(sessions.GetCurrentUserName(r))
		err = nil
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