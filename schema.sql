CREATE TABLE Utenti(
	IdUtente INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	NomeUtente CHAR(25) NOT NULL,
	CognomeUtente CHAR(25) NOT NULL,
	SessoUtente INTEGER NOT NULL,
	DataNascitaUtente integer NOT NULL,
	EmailUtente CHAR(50) NOT NULL,
	PasswordUtente VARCHAR(15) NOT NULL,
	IsCicerone BOOLEAN
);

CREATE TABLE Ciceroni(
	IdCicerone INTEGER PRIMARY KEY,
	FkUtenteCicerone INTEGER REFERENCES Utente(IdUtente),
	CodiceFiscaleCicerone CHAR(16) NOT NULL,
	TelefonoCicerone INTEGER NOT NULL,
	IbanCicerone CHAR(27) NOT NULL
);

CREATE TABLE Regioni(
	IdRegione INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	NomeRegione CHAR(30) NOT NULL
);

CREATE TABLE Province(
	IdProvincia INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkRegioneProvincia INTEGER REFERENCES Regione(IdRegione),
	NomeProvincia CHAR(25) NOT NULL
);

CREATE TABLE Citta(
	IdCitta	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkProvinciaCitta INTEGER REFERENCES Province(IdProvincia),
	NomeCitta CHAR(25) NOT NULL
);


CREATE TABLE Cap(
	IdCap INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	NumeroCap INTEGER NOT NULL
);

CREATE TABLE CapCit(
	IdCapCit INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkCittaCapCit INTEGER REFERENCES Citta(IdCitta),
	FKCapCapCit INTEGER REFERENCES Cap(IdCap),
	UNIQUE(FkCittaCapCit, FkCapCapCit)
);

CREATE TABLE Eventi(
	IdEvento INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkCiceroneEvento INTEGER REFERENCES Ciceroni(IdCicerone),
	FkCittaEvento INTEGER REFERENCES Citta(IdCitta),
	DataInizioEvento integer NOT NULL,
	DataFineEvento integer NOT NULL,
	OraInizioEvento time NOT NULL,
	OraFineEvento time NOT NULL,
	DescrizioneEvento TEXT NOT NULL,
	ItinerarioEvento TEXT NOT NULL,
	NumeroMinPart INTEGER NOT NULL,
	NumeroMaxPart INTEGER NOT NULL CHECK(NumeroMaxPart>=NumeroMinPart),
	CostoEvento REAL NOT NULL,
	IndirizzoPartenzaEvento CHAR(30) NOT NULL,
	DataScadenzaPrenotazione integer NOT NULL,
	CategoriaEvento CHAR(25) NOT NULL
);

CREATE TABLE Feedback(
	IdFeedback INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkUtenteFeedback INTEGER REFERENCES Utenti(IdUtente),
	FkEventoFeedback INTEGER REFERENCES Eventi(IdEvento),
	CommentoFeedback TEXT NOT NULL,
	DataFeedback INTEGER NOT NULL
);

CREATE TABLE Lingue(
	IdLingua INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	NomeLingua CHAR(25) NOT NULL
);

CREATE TABLE EveLin(
	IdEveLin INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkEventoEveLin INTEGER REFERENCES Eventi(IdEvento),
	FkLinguaEveLin INTEGER REFERENCES Lingue(IdLingua),
	UNIQUE(FkEventoEveLin, FkLinguaEveLin)
);

CREATE TABLE Prenotazioni(
	IdPrenotazione INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkUtentePrenotazione REFERENCES Utenti(IdUtente),
	FkEventoPrenotazione REFERENCES Eventi(IdEvento),
	DataPrenotazione INTEGER NOT NULL,
	FlagAccettazionePrenotazione BOOLEAN NOT NULL
);
