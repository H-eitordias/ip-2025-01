package main

import (
	"fmt"
)

func main() {
	var id int
	var n1, n2, n3, me float64

	fmt.Print("Número do aluno: ")
	fmt.Scanln(&id)
	fmt.Print("Notas (n1, n2, n3): ")
	fmt.Scanln(&n1, &n2, &n3)
	fmt.Print("Média dos exercícios: ")
	fmt.Scanln(&me)

	media := (n1 + 2*n2 + 3*n3 + me) / 7
	conceito := ""
	if media >= 9 {
		conceito = "A"
	} else if media >= 7.5 {
		conceito = "B"
	} else if media >= 6 {
		conceito = "C"
	} else if media >= 4 {
		conceito = "D"
	} else {
		conceito = "E"
	}

	status := "REPROVADO"
	if conceito == "A" || conceito == "B" || conceito == "C" {
		status = "APROVADO"
	}

	fmt.Println("Média:", media, "Conceito:", conceito, "-", status)
}
