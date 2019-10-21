package views

/*Holds the fetch related view handlers*/

import (
	"html/template"
	//"log"
	"net/http"
	//"time"

	//"github.com/alexanderi96/cicerone/db"
	//"github.com/alexanderi96/cicerone/sessions"
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
		//context := types.LoadContext(r)
		//username := sessions.GetCurrentUserName(r)
		//log.Println(context.user.Username)
		//currentUser, err := db.GetUserInfo(username)
		c, err := types.LoadContext(r)
		err = nil
		if err != nil {
			http.Redirect(	w, r, "/", http.StatusInternalServerError)
		} else {
			//context.CSRFToken = "abcd" I need that to utilize cookies 
			//expiration := time.Now().Add(365 * 24 * time.Hour)
			//cookie := http.Cookie{Name: "csrftoken", Value: "abcd", Expires: expiration}
			//http.SetCookie(w, &cookie)
			homeTemplate.Execute(w, c)
			//replace nil with context to load things in the html document. as far as i can see context is of type Event. there we can find a section called contentHTML, maybe used to render contents in the home template
		}
	}
}