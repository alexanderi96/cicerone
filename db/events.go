package db

import (
	"log"
	"database/sql" 
	"gitlab.com/alexanderi96/cicerone/types"
)


func CreateEvent(ciceid, idcitta, MinPart, MaxPart, costo int, dataIni, dataFine, desc, itiner, ritrovo, prenFinoAl, Categoria string) error {
	err := gQuery("insert into Evento (FkCiceroneEvento, FkCittaEvento, DataInizio, DataFine, Descrizione, Itinerario, MinPart, MaxPart, Costo, LuogoRitrovo, PrenotabileFinoAl, Categoria) values(?,?,?,?,?,?,?,?,?,?,?,?)", ciceid, idcitta, dataIni, dataFine, desc, itiner, MinPart, MaxPart, costo, ritrovo, prenFinoAl, Categoria)
	return err
}

func GetEvents() (*[]types.Evento, error) {
	log.Println("Getting events")
	var Eventi []types.Evento
	var Evento types.Evento
	var rows *sql.Rows

	basicSQL := "select * from Evento"
	rows = database.query(basicSQL)

	defer rows.Close()
	for rows.Next() {
		Evento = types.Evento{}

		err = rows.Scan(&Evento.IdEvento, &Evento.Creatore, &Evento.Citta, &Evento.DataInizio, &Evento.DataFine, &Evento.Titolo, &Evento.Descrizione, &Evento.MinPart, &Evento.MaxPart, &Evento.Costo, &Evento.Indirizzo, &Evento.DataScadenzaPren, &Evento.Categoria, &Evento.Lingue, &Evento.Prenotazioni)
		Eventi = append(Eventi, Evento)

	}

	log.Println(Eventi)

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
