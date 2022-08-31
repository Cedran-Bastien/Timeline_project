package main

import "Timeline_project/src/Jeu"

func main() {
	var jeux Jeu.Jeu = Jeu.CreerJeu(5, "donnée/timeline.txt")
	for {
		Jeu.ChoisirAction(jeux)
	}

}
