package main

import "fmt"

func StartQuizInfo() {
	fmt.Println("=== Quiz Informatique ===")
	fmt.Println("Question 1: Quelle est la capitale de la France ?")
	fmt.Println("1. Paris")
	fmt.Println("2. Londres")
	fmt.Println("3. Berlin")
	var rep int
	fmt.Print("Votre réponse : ")
	fmt.Scan(&rep)
}

/*func CheckAnswer(correctAnswer int) bool {

}

func CalculateScore(score int, total int) {

}*/
