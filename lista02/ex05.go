package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scanln(&numero)

	if numero%5 == 0 {
		fmt.Println("Divisível por 5.")
	} else {
		fmt.Println("Não é divisível por 5.")
	}
}
