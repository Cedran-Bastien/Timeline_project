package Jeu

import "Timeline_project/gestionPaquetCarte"

type Frise struct {
	frise []gestionPaquetCarte.Carte
}

func creerFrise(paquet gestionPaquetCarte.Paquet) (frise Frise) {
	var c gestionPaquetCarte.Carte = gestionPaquetCarte.PiocherCarte(paquet)

}


