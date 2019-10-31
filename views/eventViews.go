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
)

func AddEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	} 
	r.ParseForm()

	ciceid, _ := db.GetUserID(sessions.GetCurrentUser(r))
	idcitta := 1
	dataIni := r.Form.Get("iban")
	dataFine := r.Form.Get("iban")
	desc := r.Form.Get("iban")
	itiner := r.Form.Get("iban")
	MinPart, _ := strconv.Atoi(r.Form.Get("tel"))
	MaxPart, _ := strconv.Atoi(r.Form.Get("tel"))
	costo, _ := strconv.Atoi(r.Form.Get("tel"))
	ritrovo := r.Form.Get("iban")
	prenFinoAl := r.Form.Get("iban")
	Categoria := r.Form.Get("categoria")	

	err := db.CreateEvent(ciceid, idcitta, dataIni, dataFine, desc, itiner, MinPart, MaxPart, costo, ritrovo, prenFinoAl, Categoria)
	if err != nil {
		http.Error(w, "Unable to add Event", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}