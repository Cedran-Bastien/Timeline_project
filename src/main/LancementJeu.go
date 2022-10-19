package main

import (
	"Timeline_project/src/gestionPaquetCarte"
	"fmt"
)

func main() {
	var carte gestionPaquetCarte.Carte = gestionPaquetCarte.CreerCarte("bonjours:1700")
	fmt.Printf("%s", gestionPaquetCarte.ToStringCarte(carte))
	//var jeux Jeu.Jeu = Jeu.CreerJeu(5, "donnée/timeline.txt")
	//for {
	//	Jeu.ChoisirAction(jeux)
	//}

}
