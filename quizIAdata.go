package main

import "fmt"

func StartQuizIAData() {
	fmt.Println("=== Quiz IA/Data ===")
	score := 0

	fmt.Println("Question 1 : que signifie “donnée structurée” ?")
	fmt.Println("1. Une donnée organisée en lignes et colonnes")
	fmt.Println("2. Une image")
	fmt.Println("3. Un texte libre")
	if CheckAnswer(1) {
		score = score + 1
	}
	fmt.Println("Question 2: Que signifie “SQL”")
	fmt.Println("1. Super Quick Language")
	fmt.Println("2. Structured Query Language")
	fmt.Println("3. Simple Question List")
	if CheckAnswer(2) {
		score = score + 1
	}
	fmt.Println("Question 3 : Un “dataset” est :")
	fmt.Println("1. Une vidéo")
	fmt.Println("2. Un ensemble de données")
	fmt.Println("3. Une application")
	if CheckAnswer(2) {
		score = score + 1
	}
	CalculateScore(score, 3)
}
