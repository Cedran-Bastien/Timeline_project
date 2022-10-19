package test

import (
	"Timeline_project/src/gestionPaquetCarte"
	"strings"
	"testing"
)

var c gestionPaquetCarte.Carte = gestionPaquetCarte.CreerCarte("bonjour:1500")

func TestCreerCarte(t *testing.T) {

	if c.Dates != 1500 && c.Evenement != "bonjour" {
		t.Errorf("date incorrect : attendu: %d, recu: %d ", 1500, c.Dates)
	}
	if c.Evenement != "bonjour" {
		t.Errorf("evenement incorrect : attendu: %s, recu: %s", "bonjour", c.Evenement)
	}

	if c.Recto != false {
		t.Errorf("cote de la carte incorrect : attendu %t, recu %t", false, c.Recto)
	}
}

func TestRetournerCarte(t *testing.T) {
	res := gestionPaquetCarte.RetournerCarte(c)

	if res != true {
		t.Errorf("attendu : %t recu: %t", true, res)
	}
}

func TestToString(t *testing.T) {
	res := gestionPaquetCarte.ToStringCarte(c)

	if strings.Compare(res, " ??? --> bonjour") != 0 {
		t.Errorf("l'ecriture n'est aps la bonne : attendu: %s recu :%s", "??? --> bonjour", res)
	}

}
