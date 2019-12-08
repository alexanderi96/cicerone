package db

import (
	"log"
	"database/sql" 
	"gitlab.com/alexanderi96/cicerone/types"
)


func CreateEvent(e types.Evento) error {
	log.Println(e)
	err := gQuery("insert into Evento (FkCiceroneEvento, DataInizio, DataFine, Titolo, Descrizione, Itinerario, MinPart, MaxPart, Costo, LuogoRitrovo, PrenotabileFinoAl, Categoria) values(?,?,?,?,?,?,?,?,?,?,?,?)",
					e.Creatore, e.DataInizio, e.DataFine, e.Titolo,
					e.Descrizione, e.Itinerario, e.MinPart,
					e.MaxPart, e.Costo, e.Indirizzo,
					e.DataScadenzaPren, e.Categoria)
	return err
}

func GetEvents() (*[]types.Evento, error) {
	log.Println("Getting events")
	var Eventi []types.Evento
	var Evento types.Evento
	var rows *sql.Rows

	basicSQL := "select IdEvento, FkCiceroneEvento, DataInizio, DataFine, Titolo, Descrizione, Itinerario, MinPart, MaxPart, Costo, LuogoRitrovo, PrenotabileFinoAl from Evento"
	rows = database.query(basicSQL)

	defer rows.Close()
	for rows.Next() {
		Evento = types.Evento {}

		err = rows.Scan(&Evento.IdEvento, &Evento.Creatore, &Evento.DataInizio, &Evento.DataFine, &Evento.Titolo, &Evento.Descrizione, &Evento.Itinerario, &Evento.MinPart, &Evento.MaxPart, &Evento.Costo, &Evento.Indirizzo, &Evento.DataScadenzaPren)
		Eventi = append(Eventi, Evento)

	}


	return &Eventi, nil

}

func DeleteEveryEvent() error {
	
	basicSQL := "delete from Evento"
	err := gQuery(basicSQL)

	if err != nil {
		return err
	}

	return nil
}

func DeleteEvent(id int) error {
	
	basicSQL := "delete from Evento"
	err := gQuery(basicSQL)

	if err != nil {
		return err
	}

	return nil
}
