package main

import (
	"fmt"
)

func StartQuizCyber() {
	fmt.Println("=== Quiz cybersécurité ===")
	score := 0

	fmt.Println("Question 1: Que signifie “VPN” ?")
	fmt.Println("1. Virtual Private Network")
	fmt.Println("2. Very Protected Network")
	fmt.Println("3. Virtual Password Number")
	if CheckAnswer(1) {
		score = score + 1
	}
	fmt.Println("Question 2: Un antivirus sert à :")
	fmt.Println("1. Créer des mots de passe")
	fmt.Println("2. Protéger contre les logiciels malveillants")
	fmt.Println("3. Augmenter la vitesse de l’ordinateur")
	if CheckAnswer(2) {
		score = score + 1
	}
	fmt.Println("Question 3 : Que signifie “https” dans une adresse web ?")
	fmt.Println("1. Site non sécurisé")
	fmt.Println("2. Connexion chiffrée")
	fmt.Println("3. Jeu en ligne")
	if CheckAnswer(2) {
		score = score + 1
	}
	CalculateScore(score, 3)
}
