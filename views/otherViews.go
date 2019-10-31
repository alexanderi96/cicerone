package views

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	//"time"

	"gitlab.com/alexanderi96/cicerone/db"
	"gitlab.com/alexanderi96/cicerone/sessions"
	//"gitlab.com/alexanderi96/cicerone/utils"
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
	templates, err = template.ParseFiles(allFiles...)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	homeTemplate = templates.Lookup("home.html")
	loginTemplate = templates.Lookup("login.html")
	profileTemplate = templates.Lookup("userprofile.html")
}

func SignUpFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/login/", http.StatusBadRequest)
		return
	}
	r.ParseForm()

	//must add other elements
	name := r.Form.Get("name")
	surname := r.Form.Get("surname")
	gender := r.Form.Get("gender")
	birthdate := r.Form.Get("bdate") //stored as YYYY-MM-DD
	email := r.Form.Get("email")
	//must hash the passwd
	password := r.Form.Get("password")

	log.Println(name, surname, gender, email, password, birthdate)

	
	err = db.CreateUser(name, surname, gender, email, password, birthdate)
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

	uid, _ := db.GetUserID(sessions.GetCurrentUser(r))

	err := db.AddCicerone(uid, tel, iban, fcode)
	if err != nil {
		http.Error(w, "Unable to register as Cicerone", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}