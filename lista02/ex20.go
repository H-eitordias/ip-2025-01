package main

import (
	"fmt"
)

func main() {
	var preco float64
	var codigo int

	fmt.Print("Digite o preço do produto: R$ ")
	fmt.Scanln(&preco)
	fmt.Print("Digite o código da condição (1 a 4): ")
	fmt.Scanln(&codigo)

	if codigo == 1 {
		preco *= 0.90 
	} else if codigo == 2 {
		preco *= 0.95 
	} else if codigo == 4 {
		preco *= 1.10 
	} // código 3: preço normal

	fmt.Printf("Preço final: R$ %.2f\n", preco)
}
