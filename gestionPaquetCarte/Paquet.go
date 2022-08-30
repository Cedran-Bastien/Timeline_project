package gestionPaquetCarte

import (
	"bufio"
	"log"
	"math/rand"
	"os"
)

//definition d'un paquet
type Paquet struct {
	Cartes []Carte
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
		p.Cartes = append(p.Cartes, creerCarte(fileScanner.Text()))
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
	indiceCarte := rand.Intn(len(p.Cartes))
	c = p.Cartes[indiceCarte]

	//verif
	println(p.Cartes)
	println(toString(c))

	//changement tableau pour enlever carte piocher du paquet
	var nouveauPaquet []Carte
	j := 0
	for i := j; i < len(p.Cartes); i++ {
		if i == indiceCarte {
			break
		}
		nouveauPaquet = append(nouveauPaquet, p.Cartes[i])
		j++
	}
	j++
	for i := j; i < len(p.Cartes)-1; i++ {
		nouveauPaquet = append(nouveauPaquet, p.Cartes[i+1])
	}
	p.Cartes = nouveauPaquet

	//veriffication
	println(p.Cartes)

	return
}
