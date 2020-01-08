package views

import (
	//"html/template"
	//"io/ioutil"
	"log"
	"net/http"
	//"os"
	"strconv"
	//"strings"
	//"time"

	"github.com/alexanderi96/cicerone/db"
	"github.com/alexanderi96/cicerone/sessions"
	"github.com/alexanderi96/cicerone/types"
	"github.com/alexanderi96/cicerone/utils"
)

func AddEvent(w http.ResponseWriter, r *http.Request) {
	var ev types.Evento
	
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	} 
	r.ParseForm()
	ev.Creatore = db.GetUserID(sessions.GetCurrentUser(r))

	ev.DataInizio = utils.DateToUnix(r.Form.Get("dataIni"))
	ev.DataFine = utils.DateToUnix(r.Form.Get("dataFine"))
	
	ev.Titolo = r.Form.Get("title")
	ev.Descrizione = r.Form.Get("desc")
	
	//TODO: we still not have any city management query/function
	//ev.Citta, _ = strconv.Atoi(r.Form.Get("city"))
	
	ev.Itinerario = r.Form.Get("itiner")
	ev.MinPart, _ = strconv.Atoi(r.Form.Get("minP"))
	ev.MaxPart, _ = strconv.Atoi(r.Form.Get("maxP"))
	ev.Costo, _ = strconv.Atoi(r.Form.Get("cost"))
	ev.Indirizzo = r.Form.Get("meet") //must check
	ev.DataScadenzaPren = utils.DateToUnix(r.Form.Get("expDate"))
	ev.Categoria = r.Form.Get("category")	
	e = db.CreateEvent(ev)
	if e != nil {
		http.Error(w, "Unable to add Event", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func DeleteEventFunction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	r.ParseForm()

	id := r.Form.Get("id")
	log.Println(id)
	if id == "all" {
		log.Println("sus")
		err := db.DeleteEveryEvent()
		if err != nil {
			log.Println("Error deleting tasks, ", err)
			http.Redirect(w, r, "/", http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		id, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusBadRequest)
		} else {
			err = db.DeleteEventById(id)
			if err != nil {
				log.Println("Error deleting task, ", err)
			} else {
				log.Println("Event deleted")
			}
			http.Redirect(w, r, "/deleted", http.StatusFound)
		}
	}
}

//TODO: book event function
func BookEvent(w http.ResponseWriter, r *http.Request) {

}
