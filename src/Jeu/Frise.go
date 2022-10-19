package Jeu

import "Timeline_project/src/gestionPaquetCarte"

//definition du type frise
type Frise struct {
	frise gestionPaquetCarte.Paquet
}

//creer une frise chronologique
func creerFrise(paquet gestionPaquetCarte.Paquet) (frise Frise) {
	for i := 0; i < 2; i++ {
		var c gestionPaquetCarte.Carte = gestionPaquetCarte.Piocher_Carte(paquet)
		ajouterTri(frise, c)
	}
	return
}

//ajoute une carte a une position donnée d'une frise
func AjouterCarte(f Frise, c gestionPaquetCarte.Carte, pos int) {
	var tabprovisoire []gestionPaquetCarte.Carte

	for i := 0; i <= len(f.frise.Cartes); i++ {
		if i == pos {
			tabprovisoire = append(tabprovisoire, c)
		}
		if i < pos {
			tabprovisoire = append(tabprovisoire, f.frise.Cartes[i])
		}
		{
			tabprovisoire = append(tabprovisoire, f.frise.Cartes[i-1])
		}
	}
	f.frise.Cartes = tabprovisoire
}

//ajoute une carte dans l'ordre chronologique d'une frise
func ajouterTri(f Frise, c gestionPaquetCarte.Carte) {
	for i := 0; i < len(f.frise.Cartes); i++ {
		if gestionPaquetCarte.GetDates(f.frise.Cartes[i]) < gestionPaquetCarte.GetDates(c) {
			AjouterCarte(f, c, i)
		}
	}
}
