package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&numero)

	if numero > 0 {
		fmt.Println("Número positivo.")
	} else if numero < 0 {
		fmt.Println("Número negativo.")
	} else {
		fmt.Println("Número nulo.")
	}
}
