package views

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"strconv"
	"strings"

	//"github.com/alexanderi96/cicerone/db"
	//"github.com/alexanderi96/cicerone/sessions"
	//"github.com/alexanderi96/cicerone/utils"
)

//PopulateTemplates is used to parse all templates present in
//the templates folder
func PopulateTemplates() {
	var allFiles []string
	templatesDir := "./templates/"
	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		log.Println(err)
		os.Exit(1) // No point in running app if templates aren't read
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") {
			allFiles = append(allFiles, templatesDir+filename)
		}
	}

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	templates, err := template.ParseFiles(allFiles...) //added := instead of = to make it compile. must figure out
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	homeTemplate = templates.Lookup("home.html")
	//deletedTemplate = templates.Lookup("deleted.html")

	//editTemplate = templates.Lookup("edit.html")
	//searchTemplate = templates.Lookup("search.html")
	//completedTemplate = templates.Lookup("completed.html")
	loginTemplate = templates.Lookup("login.html")
	signupTemplate = templates.Lookup("signup.html")

}

func SignUpFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	r.ParseForm()

	name := r.Form.Get("name")
	surname := r.Form.Get("surname")
	username := r.Form.Get("username")
	gender := r.Form.Get("gender")
	birthdate := r.Form.Get("birthdate")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	log.Println(name, surname, username, gender, birthdate, email, password)

	/*
	err := db.CreateUser(username, password, email)
	if err != nil {
		http.Error(w, "Unable to sign user up", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/login/", 302)
	}
	*/
}