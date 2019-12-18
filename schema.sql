CREATE TABLE Utenti(
	IdUtente INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	NomeUtente CHAR(25) NOT NULL,
	CognomeUtente CHAR(25) NOT NULL,
	SessoUtente INTEGER NOT NULL,
	DataNascitaUtente integer NOT NULL,
	EmailUtente CHAR(50) NOT NULL,
	PasswordUtente VARCHAR(15) NOT NULL
);

CREATE TABLE Ciceroni(
	IdCicerone REFERENCES Utente(IdUtente),
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
	FkRegioneProvincia REFERENCES Regioni(IdRegione),
	NomeProvincia CHAR(25) NOT NULL
);

CREATE TABLE Citta(
	IdCitta	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkProvinciaCitta REFERENCES Province(IdProvincia),
	NomeCitta CHAR(25) NOT NULL
);


CREATE TABLE Cap(
	IdCap INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	NumeroCap INTEGER NOT NULL
);

CREATE TABLE CapCit(
	IdCapCit INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkCittaCapCit REFERENCES Citta(IdCitta),
	FKCapCapCit REFERENCES Cap(IdCap),
	UNIQUE(FkCittaCapCit, FkCapCapCit)
);

CREATE TABLE Eventi(
	IdEvento INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkCiceroneEvento REFERENCES Ciceroni(IdCicerone),
	FkCittaEvento REFERENCES Citta(IdCitta),
	TitoloEvento CHAR(25) NOT NULL,
	DataInizioEvento INTEGER NOT NULL,
	DataFineEvento INTEGER NOT NULL,
	DescrizioneEvento TEXT NOT NULL,
	ItinerarioEvento TEXT NOT NULL,
	NumeroMinPart INTEGER NOT NULL,
	NumeroMaxPart INTEGER NOT NULL CHECK(NumeroMaxPart>=NumeroMinPart),
	CostoEvento REAL NOT NULL,
	LuogoRitrovoEvento CHAR(30) NOT NULL,
	PrenotabileFinoAl INTEGER NOT NULL
);

CREATE TABLE Feedback(
	IdFeedback INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkUtenteFeedback REFERENCES Utenti(IdUtente),
	FkEventoFeedback REFERENCES Eventi(IdEvento),
	CommentoFeedback TEXT NOT NULL,
	DataFeedback INTEGER NOT NULL
);

CREATE TABLE Lingue(
	IdLingua INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	NomeLingua CHAR(25) NOT NULL
);

CREATE TABLE EveLin(
	IdEveLin INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkEventoEveLin REFERENCES Eventi(IdEvento),
	FkLinguaEveLin REFERENCES Lingue(IdLingua),
	UNIQUE(FkEventoEveLin, FkLinguaEveLin)
);

CREATE TABLE Prenotazioni(
	IdPrenotazione INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkUtentePrenotazione REFERENCES Utenti(IdUtente),
	FkEventoPrenotazione REFERENCES Eventi(IdEvento),
	DataPrenotazione INTEGER NOT NULL,
	FlagAccettazionePrenotazione BOOLEAN NOT NULL
);