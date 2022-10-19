package gestionPaquetCarte

import (
	"bufio"
	"math/rand"
	"os"
)

//definition d'un paquet
type Paquet struct {
	Cartes []Carte
}

//creer un paquet
//a partir d'un autre paquet
func CreerPaquetAutre(pioche Paquet, taille int) (p Paquet) {
	for i := 0; i < taille; i++ {
		PiocherCarte(pioche, p)
	}
	return p
}

//creer un paquet de carte
//a aprtir d'un fichier
func CreerPaquetFile(chemin string) (p Paquet, errOuverture error, errParcour error) {
	//ouverture du fichier
	file, err := os.Open(chemin)

	//gestion d'erreur ouverture fichier
	if err != nil {
		errOuverture = err
	}

	fileScanner := bufio.NewScanner(file)

	// remplir le paquet
	for fileScanner.Scan() {
		p.Cartes = append(p.Cartes, CreerCarte(fileScanner.Text()))
	}
	// gestion d'erreur lecture fichier
	if err := fileScanner.Err(); err != nil {
		errParcour = err
	}
	//fermetuire du fichier
	file.Close()

	return
}

// PiocherCarte pioche une carte dans un paquet
//la retourne
func PiocherCarte(p Paquet, pdest Paquet) {

	//recuperation de la carte
	indiceCarte := rand.Intn(len(p.Cartes))
	retour := p.Cartes[indiceCarte]

	//verif
	println(p.Cartes)
	println(ToStringCarte(retour))

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
	//ajout de la carte pioché dans le paquet destination
	pdest.Cartes = append(pdest.Cartes, retour)

	//veriffication
	println(p.Cartes)

}

func Piocher_Carte(p Paquet) (c Carte) {

	//recuperation de la carte
	indiceCarte := rand.Intn(len(p.Cartes))
	retour := p.Cartes[indiceCarte]

	//verif
	println(p.Cartes)
	println(ToStringCarte(retour))

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

//retourne une chaine de caractere representant le paquet
func ToStringPaquet(p Paquet) (s string) {
	for i, v := range p.Cartes {
		s += string(i) + "." + ToStringCarte(v) + "\n"
	}
	return
}
