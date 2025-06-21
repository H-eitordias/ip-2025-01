package main

import (
	"fmt"
)

func main() {
	var destino, tipo int
	var preco float64

	fmt.Print("Destino (1-Norte, 2-Nordeste, 3-Centro-Oeste, 4-Sul): ")
	fmt.Scanln(&destino)

	fmt.Print("Tipo (1-Ida e Volta, 2-Só Ida): ")
	fmt.Scanln(&tipo)

	if destino == 1 {
		if tipo == 1 {
			preco = 900
		} else {
			preco = 500
		}
	} else if destino == 2 {
		if tipo == 1 {
			preco = 650
		} else {
			preco = 350
		}
	} else if destino == 3 {
		if tipo == 1 {
			preco = 600
		} else {
			preco = 350
		}
	} else if destino == 4 {
		if tipo == 1 {
			preco = 550
		} else {
			preco = 300
		}
	} else {
		fmt.Println("Destino inválido.")
		return
	}

	fmt.Printf("Preço da passagem: R$ %.2f\n", preco)
}
