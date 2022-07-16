package gestionPaquetCarte

import (
	"log"
	"strconv"
)

//definition d'une carte
type Carte struct {
	evenement string
	dates     int
	recto     bool
}

//creer une carte a partir d'une chaine de caractere([evenement]:[date])
//toruner face verso par default
func creerCarte(s string) (c Carte) {
	var evenementAct string
	var place int = 0
	for i := place; s[i] != ':' && place < len(s); i++ {
		evenementAct += string(s[i])
		place = i
	}
	place += 1
	var dateAct string
	for i := place; s[i] != ':' && i < len(s); i++ {
		dateAct += string(s[i])
	}
	intDate, err := strconv.Atoi(dateAct)

	if err != nil {
		log.Fatalf("error during convertion string to int : %s", err)
	}
	c = Carte{evenementAct, intDate, false}
	return
}

//permet de retourner une carte
//retourne la valeur
func retournerCarte(carte Carte) bool {
	carte.recto = !carte.recto
	return carte.recto
}

//toString
func toString(c Carte) (res string) {
	if c.recto {
		res = string(c.dates) + " --> " + c.evenement
	} else {
		res = " ??? --> " + c.evenement
	}

	return
}
