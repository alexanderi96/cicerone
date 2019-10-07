package cicerone

import (
	"fmt"
	"net/http"
	"strings"
	"log"

	"github.com/alexanderi96/views"
)

func main(){

	// login-logout handlers
	http.HandleFunc("/login/", views.LoginFunc)
	http.HandleFunc("/logout/", views.RequiresLogin(views.LogoutFunc)) 
	http.HandleFunc("/signup/", SignUpFunc) 

	
	err := http.ListenAndServe(":80", nil) // listen to the port 80 for any request
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}