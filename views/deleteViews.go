package views

import (
	"log"
	"net/http"
	"strconv"
	"gitlab.com/alexanderi96/cicerone/db"
)

//MUST USE ONLY TO REFRESH DATABASE'S EVENT TABLE
func DeleteEventFunction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	log.Println("saaaaaaaaaas")

	id := r.URL.Path[len("/delete/"):]
	if id == "all" {
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
			err = db.DeleteEvent(id)
			if err != nil {
				log.Println("Error deleting task, ", err)
			} else {
				log.Println("Event deleted")
			}
			http.Redirect(w, r, "/deleted", http.StatusFound)
		}
	}
}