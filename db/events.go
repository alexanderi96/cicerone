package db

import (
	//"log"
	//"gitlab.com/alexanderi96/cicerone/db"
)


func CreateEvent(ciceid, idcitta int, dataIni, dataFine, desc, itiner string, MinPart, MaxPart, costo int, ritrovo, prenFinoAl, Categoria string) error {
	err := gQuery("insert into Evento (Cicerone, Citta, DataInizio, DataFine, Descrizione, Itinerario, MinPart, MaxPart, Costo, LuogoRitrovo, PrenotabileFinoAl, Categoria) values(?,?,?,?,?,?,?,?,?,?)", ciceid, idcitta, dataIni, dataFine, desc, itiner, MinPart, MaxPart, costo, ritrovo, prenFinoAl, Categoria)
	return err
}
