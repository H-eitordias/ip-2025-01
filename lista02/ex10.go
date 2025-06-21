package main

import (
	"fmt"
)

func main() {
	var destino, tipo int
	var preco float64

	fmt.Println("1 - Região Norte")
	fmt.Println("2 - Região Nordeste")
	fmt.Println("3 - Região Centro-Oeste")
	fmt.Println("4 - Região Sul")
	fmt.Print("Digite o destino (1 a 4): ")
	fmt.Scanln(&destino)

	fmt.Println("1 - Ida e volta")
	fmt.Println("2 - Só ida")
	fmt.Print("Inclui retorno? (1 para sim, 2 para não): ")
	fmt.Scanln(&tipo)

	if destino < 1 || destino > 4 || tipo < 1 || tipo > 2 {
		fmt.Println("Opção inválida.")
		return
	}

	switch destino {
	case 1:
		if tipo == 1 {
			preco = 900
		} else {
			preco = 500
		}
	case 2:
		if tipo == 1 {
			preco = 650
		} else {
			preco = 350
		}
	case 3:
		if tipo == 1 {
			preco = 600
		} else {
			preco = 350
		}
	case 4:
		if tipo == 1 {
			preco = 550
		} else {
			preco = 300
		}
	}

	fmt.Printf("Preço da passagem: R$ %.2f\n", preco)
}
