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

	"github.com/alexanderi96/cicerone/db"
	"github.com/alexanderi96/cicerone/sessions"
	"github.com/alexanderi96/cicerone/utils"
	"github.com/alexanderi96/cicerone/types"
)

//PopulateTemplates is used to parse all templates present in
//the templates folder
func PopulateTemplates() {
	var allFiles []string
	templatesDir := "./public/templates/"
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
	eventTemplate = templates.Lookup("event.html")
}

func SignUpFunc(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Redirect(w, r, "/login/", http.StatusBadRequest)
		return
	}
	r.ParseForm()

	var u types.User

	if r.Form.Get("cicerone") != "on" {
		u = types.Globetrotter{
			-1,
			r.Form.Get("name"),
			r.Form.Get("surname"),
			r.Form.Get("gender"),
			utils.DateToUnix(r.Form.Get("bdate")),
			r.Form.Get("email"),
			r.Form.Get("password")}
	} else {
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

func loadContext(usr string, path []string) (c types.Context, err error) {	
	
	c.Utente = db.GetUserInfo(usr)

	switch path[1] {
	case "myprofile":
		return
	case "event":
		//in this case we want to get the information about a single event
		EId, _ := strconv.Atoi(path[2])
		err = db.GetEvent(EId, c.Event)
		if err != nil {
			return c, err
		}
	default:
		ev := make(chan []types.MiniEvento)
		go db.GetEvents(ev)
		c.Events = <- ev
	}
	
	return	
}
