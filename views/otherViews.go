package views

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/alexanderi96/cicerone/db"
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

func loadContext(usr string, path []string) (c types.Context, e error) {	
	
	c.Utente = db.GetUserInfo(usr)

	switch path[1] {
	case "myprofile":
		return
	case "event":
		//in this case we want to get the information about a single event
		EId, _ := strconv.Atoi(path[2])
		c.Event, e = db.GetEventById(EId)
		if e != nil {
			return c, err
		}
	case "search":
		
	default:
		log.Println("Getting events")
		if c.Events, e = db.GetEvents(); e != nil {
			return c, e
		}
	}
	
	return	
}
