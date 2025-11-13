package main

import (
	"fmt"
)

func ShowMenu() {
	var choix int
	fmt.Println("=== GOQUIZ ===")
	fmt.Println("1. Informatique")
	fmt.Println("2. Cybersécurité")
	fmt.Println("3. Data / IA")
	fmt.Println("0. Quitter")
	fmt.Print("Choix : ")
	fmt.Scan(&choix)

	if choix == 1 {
		StartQuizInfo()
	} else if choix == 2 {
		StartQuizCyber()
	} else if choix == 3 {
		StartQuizIAData()
	} else {
		fmt.Println("Merci d'avoir joué 👋")
	}
}
