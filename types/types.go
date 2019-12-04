package types

/*
Package types is used to store the context struct which
is passed while templates are executed.

MUST EDIT
*/

//must add json things
type Utente struct {
	IdUtente		int
    Nome 		string
    Cognome 	string
    Sesso 			string
    DataNascita		string
    Email 			string
    Password 		string
    Cicerone 		*Cicerone
}

type Cicerone struct {
	IdCicerone 	int
	CodFis		string
	Tel 		int
	Iban		string
}

type Regione struct {
	IdRegione int
	NomeRegione	string
	Province	[]*Provincia //contiene gli id delle province
}

type Provincia struct {
	IdProvincia int
	NomeProvincia string
	Citta []*Citta 
}

type Citta struct {
	IdCitta int
	NomeCitta string
	Cap 	[]int
}

type Evento struct {
	IdEvento 	int
	Creatore	*Cicerone
	Citta 		*Citta
	DataInizio 	int64
	DataFine	int64
	Titolo string
	Descrizione	string
	MinPart		int
	MaxPart		int
	Costo		int
	Indirizzo	string
	DataScadenzaPren int64
	Categoria	string
	Lingue		[]*Lingua
	Prenotazioni []*Prenotazioni
}

type Lingua struct {
	IdLingua int
	NomeLingua string
}

type Prenotazioni struct {
	IdPrenotazione	int
	Utente 			int
	DataPrenotazione int64
	Accettazione	bool
}


//Context is the struct passed to templates
type Context struct {
	Events      *[]Evento 	//using
	Utente 		*Utente 		//using
	CSRFToken  	string
	Referer    	string
	IsCicerone	bool
	City 		*Citta
	Category 	*Category
}

type Category struct {
	IdCat 	int
	Nome 	string
}
