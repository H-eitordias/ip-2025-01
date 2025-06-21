package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Print("Digite um número inteiro de 3 dígitos: ")
	fmt.Scanln(&numero)

	if numero >= 100 && numero <= 999 {
		dezenas := (numero / 10) % 10
		fmt.Printf("Algarismo das dezenas: %d\n", dezenas)
	} else {
		fmt.Println("Número inválido. Deve ter exatamente 3 dígitos.")
	}
}
