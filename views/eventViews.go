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
	dataIni := r.Form.Get("dataIni") //must check
	dataFine := r.Form.Get("dataFine") //must check
	//title := r.Form.Get("title")
	desc := r.Form.Get("desc")
	city, _ := strconv.Atoi(r.Form.Get("city")) //must check
	itiner := r.Form.Get("itiner")
	MinPart, _ := strconv.Atoi(r.Form.Get("minP"))
	MaxPart, _ := strconv.Atoi(r.Form.Get("maxP"))
	costo, _ := strconv.Atoi(r.Form.Get("cost"))
	ritrovo := r.Form.Get("meet") //must check
	prenFinoAl := r.Form.Get("expDate")
	Categoria := r.Form.Get("category")	

	err := db.CreateEvent(ciceid, city, MinPart, MaxPart, costo, dataIni, dataFine, desc, itiner, ritrovo, prenFinoAl, Categoria)
	if err != nil {
		http.Error(w, "Unable to add Event", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}