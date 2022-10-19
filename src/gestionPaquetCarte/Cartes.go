package gestionPaquetCarte

import (
	"log"
	"strconv"
)

//definition d'une carte
type Carte struct {
	Evenement string
	Dates     int
	Recto     bool
}

//creer une carte a partir d'une chaine de caractere([evenement]:[date])
//toruner face verso par default
func CreerCarte(s string) (c Carte) {
	var evenementAct string
	var place int = 0
	for {
		if s[place] == ':' {
			break
		}
		evenementAct += string(s[place])
		place++
	}
	place += 1
	var dateAct string
	for i := place; i < len(s); i++ {
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
func RetournerCarte(carte Carte) bool {
	carte.Recto = !carte.Recto
	return carte.Recto
}

//toString
func ToStringCarte(c Carte) (res string) {
	if c.Recto {
		res = string(c.Dates) + " --> " + c.Evenement
	} else {
		res = " ??? --> " + c.Evenement
	}

	return
}

func GetDates(c Carte) int {
	return c.Dates
}
