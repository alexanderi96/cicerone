package db

import (
	"log"
	"database/sql" 
	"github.com/alexanderi96/cicerone/types"
)


func CreateEvent(e types.Evento) error {
	log.Println(e)
	err := gQuery("insert into Eventi(FkCiceroneEvento, DataInizioEvento, DataFineEvento, TitoloEvento, DescrizioneEvento, ItinerarioEvento, NumeroMinPart, NumeroMaxPart, CostoEvento, IndirizzoPartenzaEvento, DataScadenzaPrenotazione, CategoriaEvento) values(?,?,?,?,?,?,?,?,?,?,?,?)",
					e.Creatore, e.DataInizio, e.DataFine, e.Titolo,
					e.Descrizione, e.Itinerario, e.MinPart,
					e.MaxPart, e.Costo, e.Indirizzo,
					e.DataScadenzaPren, e.Categoria)
	return err
}

func GetEvents(E chan []types.MiniEvento) {
	log.Println("Getting events")
	var Eventi []types.MiniEvento
	var Evento types.MiniEvento
	var rows *sql.Rows

	basicSQL := "select IdEvento, TitoloEvento, DescrizioneEvento from Eventi"
	rows = database.query(basicSQL)

	defer rows.Close()
	for rows.Next() {
		Evento = types.MiniEvento {}
		err = rows.Scan(&Evento.IdEvento, &Evento.Titolo, &Evento.Descrizione)
		if err != nil {
			log.Println(err)
		}
		Eventi = append(Eventi, Evento)
	}
	E <- Eventi

}

func GetEvent(id int, Evento types.Evento) error {
	log.Println("Getting Event")
	var rows *sql.Rows

	basicSQL := "select IdEvento, FkCiceroneEvento, DataInizioEvento, DataFineEvento, TitoloEvento, DescrizioneEvento, ItinerarioEvento, NumeroMinPart, NumeroMaxPart, CostoEvento, IndirizzoPartenzaEvento, DataScadenzaPrenotazione from Eventi where IdEvento = ?"
	rows = database.query(basicSQL, id)

	defer rows.Close()
	if rows.Next() {

		err = rows.Scan(&Evento.IdEvento, &Evento.Creatore, &Evento.DataInizio, &Evento.DataFine, &Evento.Titolo, &Evento.Descrizione, &Evento.Itinerario, &Evento.MinPart, &Evento.MaxPart, &Evento.Costo, &Evento.Indirizzo, &Evento.DataScadenzaPren)
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteEveryEvent() error {
	
	basicSQL := "delete from Eventi"
	err := gQuery(basicSQL)

	if err != nil {
		return err
	}

	return nil
}

func DeleteEvent(id int) error {
	
	basicSQL := "delete from Eventi"
	err := gQuery(basicSQL)

	if err != nil {
		return err
	}

	return nil
}
