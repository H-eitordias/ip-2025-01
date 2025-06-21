package main

import (
	"fmt"
)

func main() {
	var dia int
	var preco float64
	var tipo string

	fmt.Print("Digite o dia da semana (1-dom a 7-sáb): ")
	fmt.Scanln(&dia)
	fmt.Print("Preço normal do DVD: R$ ")
	fmt.Scanln(&preco)
	fmt.Print("Tipo do DVD (comum ou lancamento): ")

	if dia == 2 || dia == 3 || dia == 5 {
		preco *= 0.60
	}

	if tipo == "lancamento" {
		preco *= 1.15
	}

	fmt.Printf("Preço final: R$ %.2f\n", preco)
}
