CREATE TABLE Utenti(
	IdUtente INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	NomeUtente CHAR(25) NOT NULL,
	CognomeUtente CHAR(25) NOT NULL,
	SessoUtente INTEGER NOT NULL,
	DataNascitaUtente DATE NOT NULL,
	EmailUtente CHAR(50) NOT NULL,
	PasswordUtente VARCHAR(15) NOT NULL
);

CREATE TABLE Ciceroni(
	IdCicerone INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkUtenteCicerone INTEGER NOT NULL,
	CodiceFiscaleCicerone CHAR(16) NOT NULL,
	TelefonoCicerone INTEGER NOT NULL,
	IbanCicerone CHAR(27) NOT NULL,
	FOREIGN KEY(FkUtenteCicerone) REFERENCES Utenti(IdUtente)
);

CREATE TABLE Regioni(
	IdRegione INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	NomeRegione CHAR(30) NOT NULL
);

CREATE TABLE Province(
	IdProvincia INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkRegioneProvincia INTEGER NOT NULL,
	NomeProvincia CHAR(25) NOT NULL
);

CREATE TABLE Citta(
	IdCitta	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkProvinciaCitta INTEGER NOT NULL,
	NomeCitta CHAR(25) NOT NULL,
	FOREIGN KEY(FkProvinciaCitta) REFERENCES Province(IdProvincia)
);


CREATE TABLE Cap(
	IdCap INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	NumeroCap INTEGER NOT NULL
);

CREATE TABLE CapCit(
	IdCapCit INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkCittaCapCit INTEGER NOT NULL,
	FKCapCapCit INTEGER NOT NULL,
	FOREIGN KEY(FkCittaCapCit) REFERENCES Citta(IdCitta),
	FOREIGN KEY(FKCapCapCit) REFERENCES Cap(IdCap),
	UNIQUE(FkCittaCapCit, FkCapCapCit)
);

CREATE TABLE Eventi(
	IdEvento INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkCiceroneEvento INTEGER NOT NULL,
	FkCittaEvento INTEGER NOT NULL,
	DataInizioEvento DATE NOT NULL,
	DataFineEvento DATE NOT NULL,
	OraInizioEvento TIME NOT NULL,
	OraFineEvento TIME NOT NULL,
	DescrizioneEvento TEXT NOT NULL,
	ItinerarioEvento TEXT NOT NULL,
	NumeroMinimoPartecipantiEvento INTEGER NOT NULL,
	NumeroMassimoPartecipantiEvento	INTEGER NOT NULL CHECK(NumeroMassimoPartecipantiEvento>=NumeroMinimoPartecipantiEvento),
	CostoEvento REAL NOT NULL,
	IndirizzoPartenzaEvento CHAR(30) NOT NULL,
	DataScadenzaPrenotazioneEvento DATE NOT NULL,
	CategoriaEvento CHAR(25) NOT NULL,
	FOREIGN KEY(FkCittaEvento) REFERENCES Citta(IdCitta),
	FOREIGN KEY(FkCiceroneEvento) REFERENCES Ciceroni(IdCicerone)
);

CREATE TABLE Lingue(
	IdLingua INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	NomeLingua INTEGER NOT NULL
);

CREATE TABLE EveLin(
	IdEveLin INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkEventoEveLin INTEGER NOT NULL,
	FkLinguaEveLin INTEGER NOT NULL,
	FOREIGN KEY(IdEventoEveLin) REFERENCES Eventi(IdEvento),
	UNIQUE(FkEventoEveLin, FkLinguaEveLin)
);

CREATE TABLE Prenotazioni(
	IdPrenotazione INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FkUtentePrenotazione INTEGER NOT NULL,
	FkCiceronePrenotazione INTEGER NOT NULL,
	FkEventoPrenotazione INTEGER NOT NULL,
	DataPrenotazione DATE NOT NULL,
	FlagAccettazionePrenotazione BOOLEAN NOT NULL,
	FOREIGN KEY(FkUtentePrenotazione) REFERENCES Utenti(IdUtente),
	FOREIGN KEY(FkCiceronePrenotazione) REFERENCES Ciceroni(IdCicerone),
	FOREIGN KEY(FkEventoPrenotazione) REFERENCES Eventi(IdEvento)
);