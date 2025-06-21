package main

import (
	"fmt"
)

func main() {
	var idade int
	fmt.Print("Digite a idade: ")
	fmt.Scanln(&idade)

	if idade < 16 {
		fmt.Println("Não-eleitor")
	} else if idade >= 18 && idade <= 65 {
		fmt.Println("Eleitor obrigatório")
	} else {
		fmt.Println("Eleitor facultativo")
	}
}
