-- phpMyAdmin SQL Dump
-- version 4.9.0.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Creato il: Ott 15, 2019 alle 17:24
-- Versione del server: 10.4.6-MariaDB
-- Versione PHP: 7.3.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `cicerone`
--

-- --------------------------------------------------------

--
-- Struttura della tabella `attivita`
--

CREATE TABLE `attivita` (
  `IdAttivita` int(2) UNSIGNED NOT NULL,
  `FkCiceroneAttivita` int(3) UNSIGNED NOT NULL,
  `NomeAttivita` char(30) COLLATE utf8_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------------------------

--
-- Struttura della tabella `cap`
--

CREATE TABLE `cap` (
  `IdCap` int(2) UNSIGNED NOT NULL,
  `NumeroCap` int(5) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------------------------

--
-- Struttura della tabella `capcit`
--

CREATE TABLE `capcit` (
  `IdCapCit` int(2) UNSIGNED NOT NULL,
  `FkCittaCapCit` int(2) UNSIGNED NOT NULL,
  `FkCapCapCit` int(2) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------------------------

--
-- Struttura della tabella `ciceroni`
--

CREATE TABLE `ciceroni` (
  `IdCicerone` int(3) UNSIGNED NOT NULL,
  `FkUtenteCicerone` int(3) UNSIGNED NOT NULL,
  `CodiceFiscaleCicerone` char(16) COLLATE utf8_bin NOT NULL,
  `TelefonoCicerone` int(15) UNSIGNED NOT NULL,
  `IbanCicerone` char(27) COLLATE utf8_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------------------------

--
-- Struttura della tabella `citta`
--

CREATE TABLE `citta` (
  `IdCitta` int(2) UNSIGNED NOT NULL,
  `FkProvinciaCitta` int(2) UNSIGNED NOT NULL,
  `NomeCitta` char(25) COLLATE utf8_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------------------------

--
-- Struttura della tabella `eventi`
--

CREATE TABLE `eventi` (
  `IdEvento` int(4) UNSIGNED NOT NULL,
  `FkCiceroneEvento` int(3) UNSIGNED NOT NULL,
  `FkCapCitEvento` int(2) UNSIGNED NOT NULL,
  `DataInizioEvento` date NOT NULL,
  `DataFineEvento` date NOT NULL,
  `OraInizioEvento` time NOT NULL,
  `OraFineEvento` time NOT NULL,
  `DescrizioneEvento` text COLLATE utf8_bin NOT NULL,
  `ItinerarioEvento` text COLLATE utf8_bin NOT NULL,
  `NumeroMinimoPartecipantiEvento` int(1) UNSIGNED NOT NULL,
  `NumeroMassimoPartecipantiEvento` int(1) UNSIGNED NOT NULL,
  `CostoEvento` decimal(4,2) UNSIGNED NOT NULL,
  `IndirizzoPartenzaEvento` char(30) COLLATE utf8_bin NOT NULL,
  `DataScadenzaEvento` date NOT NULL,
  `TipologiaEvento` char(25) COLLATE utf8_bin NOT NULL
) ;

-- --------------------------------------------------------

--
-- Struttura della tabella `preatt`
--

CREATE TABLE `preatt` (
  `IdPreAtt` int(2) UNSIGNED NOT NULL,
  `FkPrenotazionePreAtt` int(2) UNSIGNED NOT NULL,
  `FkAttivitaPreAtt` int(2) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------------------------

--
-- Struttura della tabella `prenotazioni`
--

CREATE TABLE `prenotazioni` (
  `IdPrenotazione` int(2) UNSIGNED NOT NULL,
  `FkUtentePrenotazione` int(3) UNSIGNED NOT NULL,
  `FkEventoPrenotazione` int(4) UNSIGNED NOT NULL,
  `DataPrenotazione` date NOT NULL,
  `FlagAccettazionePrenotazione` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------------------------

--
-- Struttura della tabella `province`
--

CREATE TABLE `province` (
  `IdProvincia` int(2) UNSIGNED NOT NULL,
  `FkRegioneProvincia` int(2) UNSIGNED NOT NULL,
  `NomeProvincia` char(25) COLLATE utf8_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------------------------

--
-- Struttura della tabella `regioni`
--

CREATE TABLE `regioni` (
  `IdRegione` int(2) UNSIGNED NOT NULL,
  `NomeRegione` char(30) COLLATE utf8_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- --------------------------------------------------------

--
-- Struttura della tabella `utenti`
--

CREATE TABLE `utenti` (
  `IdUtente` int(3) UNSIGNED NOT NULL,
  `NomeUtente` char(25) COLLATE utf8_bin NOT NULL,
  `CognomeUtente` char(25) COLLATE utf8_bin NOT NULL,
  `SessoUtente` enum('Uomo','Donna','Altro') COLLATE utf8_bin NOT NULL,
  `DataNascitaUtente` date NOT NULL,
  `EmailUtente` char(50) COLLATE utf8_bin NOT NULL,
  `PasswordUtente` varchar(15) COLLATE utf8_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

--
-- Indici per le tabelle scaricate
--

--
-- Indici per le tabelle `attivita`
--
ALTER TABLE `attivita`
  ADD PRIMARY KEY (`IdAttivita`),
  ADD KEY `FkCiceroneAttivita` (`FkCiceroneAttivita`);

--
-- Indici per le tabelle `cap`
--
ALTER TABLE `cap`
  ADD PRIMARY KEY (`IdCap`);

--
-- Indici per le tabelle `capcit`
--
ALTER TABLE `capcit`
  ADD PRIMARY KEY (`IdCapCit`),
  ADD UNIQUE KEY `FkCittaCapCit` (`FkCittaCapCit`,`FkCapCapCit`),
  ADD KEY `FkCapCapCit` (`FkCapCapCit`);

--
-- Indici per le tabelle `ciceroni`
--
ALTER TABLE `ciceroni`
  ADD PRIMARY KEY (`IdCicerone`),
  ADD KEY `FkUtenteCicerone` (`FkUtenteCicerone`);

--
-- Indici per le tabelle `citta`
--
ALTER TABLE `citta`
  ADD PRIMARY KEY (`IdCitta`),
  ADD KEY `FkProvinciaCitta` (`FkProvinciaCitta`);

--
-- Indici per le tabelle `eventi`
--
ALTER TABLE `eventi`
  ADD PRIMARY KEY (`IdEvento`),
  ADD KEY `FkCiceroneEvento` (`FkCiceroneEvento`),
  ADD KEY `FkCapCitEvento` (`FkCapCitEvento`);

--
-- Indici per le tabelle `preatt`
--
ALTER TABLE `preatt`
  ADD PRIMARY KEY (`IdPreAtt`),
  ADD UNIQUE KEY `FkPrenotazionePreAtt` (`FkPrenotazionePreAtt`,`FkAttivitaPreAtt`),
  ADD KEY `FkAttivitaPreAtt` (`FkAttivitaPreAtt`);

--
-- Indici per le tabelle `prenotazioni`
--
ALTER TABLE `prenotazioni`
  ADD PRIMARY KEY (`IdPrenotazione`),
  ADD KEY `FkUtentePrenotazione` (`FkUtentePrenotazione`),
  ADD KEY `FkEventoPrenotazione` (`FkEventoPrenotazione`);

--
-- Indici per le tabelle `province`
--
ALTER TABLE `province`
  ADD PRIMARY KEY (`IdProvincia`),
  ADD KEY `FkRegioneProvincia` (`FkRegioneProvincia`);

--
-- Indici per le tabelle `regioni`
--
ALTER TABLE `regioni`
  ADD PRIMARY KEY (`IdRegione`);

--
-- Indici per le tabelle `utenti`
--
ALTER TABLE `utenti`
  ADD PRIMARY KEY (`IdUtente`);

--
-- AUTO_INCREMENT per le tabelle scaricate
--

--
-- AUTO_INCREMENT per la tabella `attivita`
--
ALTER TABLE `attivita`
  MODIFY `IdAttivita` int(2) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT per la tabella `cap`
--
ALTER TABLE `cap`
  MODIFY `IdCap` int(2) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT per la tabella `capcit`
--
ALTER TABLE `capcit`
  MODIFY `IdCapCit` int(2) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT per la tabella `ciceroni`
--
ALTER TABLE `ciceroni`
  MODIFY `IdCicerone` int(3) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT per la tabella `citta`
--
ALTER TABLE `citta`
  MODIFY `IdCitta` int(2) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT per la tabella `eventi`
--
ALTER TABLE `eventi`
  MODIFY `IdEvento` int(4) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT per la tabella `preatt`
--
ALTER TABLE `preatt`
  MODIFY `IdPreAtt` int(2) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT per la tabella `prenotazioni`
--
ALTER TABLE `prenotazioni`
  MODIFY `IdPrenotazione` int(2) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT per la tabella `province`
--
ALTER TABLE `province`
  MODIFY `IdProvincia` int(2) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT per la tabella `regioni`
--
ALTER TABLE `regioni`
  MODIFY `IdRegione` int(2) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT per la tabella `utenti`
--
ALTER TABLE `utenti`
  MODIFY `IdUtente` int(3) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- Limiti per le tabelle scaricate
--

--
-- Limiti per la tabella `attivita`
--
ALTER TABLE `attivita`
  ADD CONSTRAINT `attivita_ibfk_1` FOREIGN KEY (`FkCiceroneAttivita`) REFERENCES `ciceroni` (`IdCicerone`);

--
-- Limiti per la tabella `capcit`
--
ALTER TABLE `capcit`
  ADD CONSTRAINT `capcit_ibfk_1` FOREIGN KEY (`FkCittaCapCit`) REFERENCES `citta` (`IdCitta`),
  ADD CONSTRAINT `capcit_ibfk_2` FOREIGN KEY (`FkCapCapCit`) REFERENCES `cap` (`IdCap`);

--
-- Limiti per la tabella `ciceroni`
--
ALTER TABLE `ciceroni`
  ADD CONSTRAINT `ciceroni_ibfk_1` FOREIGN KEY (`FkUtenteCicerone`) REFERENCES `utenti` (`IdUtente`);

--
-- Limiti per la tabella `citta`
--
ALTER TABLE `citta`
  ADD CONSTRAINT `citta_ibfk_1` FOREIGN KEY (`FkProvinciaCitta`) REFERENCES `province` (`IdProvincia`);

--
-- Limiti per la tabella `eventi`
--
ALTER TABLE `eventi`
  ADD CONSTRAINT `eventi_ibfk_1` FOREIGN KEY (`FkCiceroneEvento`) REFERENCES `ciceroni` (`IdCicerone`),
  ADD CONSTRAINT `eventi_ibfk_2` FOREIGN KEY (`FkCapCitEvento`) REFERENCES `capcit` (`IdCapCit`);

--
-- Limiti per la tabella `preatt`
--
ALTER TABLE `preatt`
  ADD CONSTRAINT `preatt_ibfk_1` FOREIGN KEY (`FkPrenotazionePreAtt`) REFERENCES `prenotazioni` (`IdPrenotazione`),
  ADD CONSTRAINT `preatt_ibfk_2` FOREIGN KEY (`FkAttivitaPreAtt`) REFERENCES `attivita` (`IdAttivita`);

--
-- Limiti per la tabella `prenotazioni`
--
ALTER TABLE `prenotazioni`
  ADD CONSTRAINT `prenotazioni_ibfk_1` FOREIGN KEY (`FkUtentePrenotazione`) REFERENCES `utenti` (`IdUtente`),
  ADD CONSTRAINT `prenotazioni_ibfk_2` FOREIGN KEY (`FkEventoPrenotazione`) REFERENCES `eventi` (`IdEvento`);

--
-- Limiti per la tabella `province`
--
ALTER TABLE `province`
  ADD CONSTRAINT `province_ibfk_1` FOREIGN KEY (`FkRegioneProvincia`) REFERENCES `regioni` (`IdRegione`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
