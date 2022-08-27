package gestionPaquetCarte

import (
	"bufio"
	"log"
	"math/rand"
	"os"
)

//definition d'un paquet
type Paquet struct {
	cartes []Carte
}

//creer un paquet de carte
//a aprtir d'un fichier
func creerPaquet(chemin string) (p Paquet) {
	//ouverture du fichier
	file, err := os.Open(chemin)

	//gestion d'erreur ouverture fichier
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	// remplir le paquet
	for fileScanner.Scan() {
		p.cartes = append(p.cartes, creerCarte(fileScanner.Text()))
	}
	// gestion d'erreur lecture fichier
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	//fermetuire du fichier
	file.Close()

	return
}

// PiocherCarte pioche une carte dans un paquet
//la retourne
func PiocherCarte(p Paquet) (c Carte) {

	//recuperation de la carte
	indiceCarte := rand.Intn(len(p.cartes))
	c = p.cartes[indiceCarte]

	//verif
	println(p.cartes)
	println(toString(c))

	//changement tableau pour enlever carte piocher du paquet
	var nouveauPaquet []Carte
	j := 0
	for i := j; i < len(p.cartes); i++ {
		if i == indiceCarte {
			break
		}
		nouveauPaquet = append(nouveauPaquet, p.cartes[i])
		j++
	}
	j++
	for i := j; i < len(p.cartes)-1; i++ {
		nouveauPaquet = append(nouveauPaquet, p.cartes[i+1])
	}
	p.cartes = nouveauPaquet

	//veriffication
	println(p.cartes)

	return
}
