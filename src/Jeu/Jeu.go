package Jeu

import (
	"Timeline_project/src/gestionPaquetCarte"
	"fmt"
	"log"
)

//definition du type jeu
type Jeu struct {
	frise  Frise
	paquet gestionPaquetCarte.Paquet
	pioche gestionPaquetCarte.Paquet
}

func CreerJeu(nbCarteJoueur int, chemin string) (j Jeu) {
	j.pioche = gestionPaquetCarte.CreerPaquetFile(chemin)
	j.frise = creerFrise(j.pioche)
	j.paquet = gestionPaquetCarte.CreerPaquetAutre(j.pioche, nbCarteJoueur)

	return
}

func ChoisirAction(jeu Jeu) (c gestionPaquetCarte.Carte, position int) {
	fmt.Printf("", gestionPaquetCarte.ToString(jeu.frise.frise))
	fmt.Printf("", gestionPaquetCarte.ToString(jeu.paquet))
	fmt.Printf("")
	position, err := fmt.Scanf("", "quel carte voulez vous placer(indice)\nreponse :")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	c = jeu.paquet.Cartes[position]
	fmt.Printf("l'indice de la carte est : %d", position)
	return
}
