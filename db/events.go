package db

import (
	"log"
	"database/sql" 
	"github.com/alexanderi96/cicerone/types"
)


func CreateEvent(e types.Evento) error {
	return gQuery("insert into Evento (FkCiceroneEvento, DataInizio, DataFine, Titolo, Descrizione, Itinerario, MinPart, MaxPart, Costo, LuogoRitrovo, PrenotabileFinoAl, Categoria) values(?,?,?,?,?,?,?,?,?,?,?,?)",
					e.Creatore, e.DataInizio, e.DataFine, e.Titolo,
					e.Descrizione, e.Itinerario, e.MinPart,
					e.MaxPart, e.Costo, e.Indirizzo,
					e.DataScadenzaPren, e.Categoria)
}

func GetEvents() (E []types.MiniEvento, e error) {
	var Evento types.MiniEvento
	var rows *sql.Rows

	basicSQL := "select IdEvento, TitoloEvento, DescrizioneEvento from Eventi"
	rows = database.query(basicSQL)

	defer rows.Close()
	for rows.Next() {
		Evento = types.MiniEvento {}
		err = rows.Scan(&Evento.IdEvento, &Evento.Titolo, &Evento.Descrizione)
		if err != nil {
			return E, e
		}
		E = append(E, Evento)
	}
	return E, nil

}

func GetEventById(id int) (Evento types.Evento, e error) {
	log.Println("Getting Event")
	var rows *sql.Rows

	basicSQL := "select IdEvento, FkCiceroneEvento, DataInizioEvento, DataFineEvento, TitoloEvento, DescrizioneEvento, ItinerarioEvento, NumeroMinPart, NumeroMaxPart, CostoEvento, IndirizzoPartenzaEvento, DataScadenzaPrenotazione from Eventi where IdEvento = ?"
	rows = database.query(basicSQL, id)

	defer rows.Close()
	if rows.Next() {

		e = rows.Scan(&Evento.IdEvento, &Evento.Creatore, &Evento.DataInizio, &Evento.DataFine, &Evento.Titolo, &Evento.Descrizione, &Evento.Itinerario, &Evento.MinPart, &Evento.MaxPart, &Evento.Costo, &Evento.Indirizzo, &Evento.DataScadenzaPren)
		if e != nil {
			Evento = types.Evento {}
			return Evento, e
		}
	}

	return Evento, nil
}

func DeleteEveryEvent() error {
	basicSQL := "delete from Evento"
	return gQuery(basicSQL)
}

func DeleteEventById(id int) (e error) {
	basicSQL := "delete from Evento where IdEvento = ?"
	e = gQuery(basicSQL, id)
	//TODO: notify every interested user that this event was deleted 
	return
}
