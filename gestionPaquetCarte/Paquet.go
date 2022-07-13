package gestionPaquetCarte

import (
	"bufio"
	"log"
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

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	// remplir le paquet
	for fileScanner.Scan() {
		p.cartes = append(p.cartes, creerCarte(fileScanner.Text()))
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	//fermetuire du fichier
	file.Close()

	return
}

