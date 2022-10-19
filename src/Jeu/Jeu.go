package Jeu

import (
	"Timeline_project/src/gestionPaquetCarte"
	"fmt"
)

//definition du type jeu
type Jeu struct {
	friseJ  Frise
	paquetJ gestionPaquetCarte.Paquet
	pioche  gestionPaquetCarte.Paquet
}

func CreerJeu(nbCarteJoueur int, chemin string) (j Jeu, errChemin error, errStructureFile error) {
	j.pioche, errChemin, errStructureFile = gestionPaquetCarte.CreerPaquetFile(chemin)
	j.friseJ = creerFrise(j.pioche)
	j.paquetJ = gestionPaquetCarte.CreerPaquetAutre(j.pioche, nbCarteJoueur)
	return
}

func ChoisirAction(jeu Jeu) (c gestionPaquetCarte.Carte, position int) {
	fmt.Printf("=====================")
	fmt.Printf("FRISE")
	fmt.Printf("=====================")
	fmt.Printf("", gestionPaquetCarte.ToStringPaquet(jeu.friseJ.frise))
	fmt.Printf("=====================")
	fmt.Printf("VOTRE MAIN :")
	fmt.Printf("=====================")
	fmt.Printf("", gestionPaquetCarte.ToStringPaquet(jeu.paquetJ))
	fmt.Printf("")
	choixCarte := false
	for choixCarte {
		fmt.Printf("quel carte voulez vous placer(indice)\nreponse :")
		_, err := fmt.Scanf("%d", position)
		if err != nil {
			fmt.Printf("error : l'entrer n'est pas un nombre ")
		} else if position > len(jeu.paquetJ.Cartes) {
			fmt.Printf("error : le nombre rentrer ne fait pas parti du paquet ")
		} else {
			choixCarte = true
		}
	}

	c = jeu.paquetJ.Cartes[position]

	choixFrise := false
	for choixFrise {

		fmt.Printf("A quel position voulez vous placer la carte %s\nreponse :", c)
		_, err := fmt.Scanf("%d", position)
		if err != nil {
			fmt.Printf("error : l'entrer n'est pas un nombre ")
		} else if position > len(jeu.friseJ.frise.Cartes) {
			fmt.Printf("error : le nombre rentrer ne fait pas parti du paquet ")
		} else {
			choixCarte = true
		}
	}

	return
}

func PlacerCarte(j Jeu, c gestionPaquetCarte.Carte, position int) (b bool) {
	b = false
	if j.friseJ.frise.Cartes[position-1].Dates < c.Dates && j.friseJ.frise.Cartes[position].Dates > c.Dates {
		AjouterCarte(j.friseJ, c, position)
		gestionPaquetCarte.PiocherCarte(j.pioche, j.paquetJ)
		b = true
		fmt.Printf("la carte a bien ete placé")
	} else {
		fmt.Printf("Carte non placer : mauvaise position ")
	}
	return
}
