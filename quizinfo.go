package main

import "fmt"

func StartQuizInfo() {

	fmt.Println("=== Quiz Informatique ===")
	score := 0

	fmt.Println("Question 1: À quoi sert le “HTML” dans un site web ?")
	fmt.Println("1. À sécuriser les informations des utilisateurs")
	fmt.Println("2. À augmenter la vitesse de chargement des sites")
	fmt.Println("3. À écrire le contenu et la structure des pages")
	if CheckAnswer(3) {
		score = score + 1
	}

	fmt.Println("Question 2: Quel langage est principalement utilisé pour le développement d'applications Android ?")
	fmt.Println("1. Java")
	fmt.Println("2. Swift")
	fmt.Println("3. Ruby")
	if CheckAnswer(1) {
		score = score + 1
	}
	fmt.Println("Question 3: Qu'est-ce qu'un 'serveur' dans le contexte de l'informatique ?")
	fmt.Println("1. Un type de logiciel malveillant")
	fmt.Println("2. Un programme qui gère les ressources réseau et fournit des services aux autres ordinateurs")
	fmt.Println("3. Un périphérique de stockage externe")
	if CheckAnswer(2) {
		score = score + 1
	}
	CalculateScore(score, 3)
}

func CheckAnswer(correctAnswer int) bool {
	var userAnswer int
	fmt.Print("Votre réponse : ")
	fmt.Scan(&userAnswer)
	if userAnswer == correctAnswer {
		fmt.Println("Bonne réponse !")
		return true
	} else {
		fmt.Println("Mauvaise réponse.")
		return false
	}
}

func CalculateScore(score int, total int) {
	fmt.Printf("Votre score final est %d sur %d.\n", score, total)
	if score == total {
		fmt.Println("Félicitations ! Vous avez obtenu un score parfait ! 🎉")
	} else if score >= total/2 {
		fmt.Println("Bon travail ! Vous avez réussi le quiz. 👍")
	} else {
		fmt.Println("Ne vous découragez pas, essayez à nouveau pour vous améliorer ! 💪")
	}
}
