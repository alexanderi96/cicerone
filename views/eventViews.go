package views

import (
	//"html/template"
	//"io/ioutil"
	//"log"
	"net/http"
	//"os"
	"strconv"
	//"strings"
	//"time"

	"gitlab.com/alexanderi96/cicerone/db"
	"gitlab.com/alexanderi96/cicerone/sessions"
	"gitlab.com/alexanderi96/cicerone/types"
	"gitlab.com/alexanderi96/cicerone/utils"
)

func AddEvent(w http.ResponseWriter, r *http.Request) {
	var e types.Evento
	
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	} 
	r.ParseForm()
	e.Creatore, _ = db.GetUserID(sessions.GetCurrentUser(r))

	e.DataInizio = utils.DateToUnix(r.Form.Get("dataIni"))
	e.DataFine = utils.DateToUnix(r.Form.Get("dataFine"))
	
	e.Titolo = r.Form.Get("title")
	e.Descrizione = r.Form.Get("desc")
	
	//TODO: we still not have any city management query/function
	//e.Citta, _ = strconv.Atoi(r.Form.Get("city"))
	
	e.Itinerario = r.Form.Get("itiner")
	e.MinPart, _ = strconv.Atoi(r.Form.Get("minP"))
	e.MaxPart, _ = strconv.Atoi(r.Form.Get("maxP"))
	e.Costo, _ = strconv.Atoi(r.Form.Get("cost"))
	e.Indirizzo = r.Form.Get("meet") //must check
	e.DataScadenzaPren = utils.DateToUnix(r.Form.Get("expDate"))
	e.Categoria = r.Form.Get("category")	
	err = db.CreateEvent(e)
	if err != nil {
		http.Error(w, "Unable to add Event", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
