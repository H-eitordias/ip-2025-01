package main

import (
	"fmt"
)

func main() {
	var precoBase float64
	var ar, pintura, vidro, direcao string
	var precoFinal float64

	fmt.Print("Digite o preço de fábrica do carro: R$ ")
	fmt.Scanln(&precoBase)

	precoFinal = precoBase

	fmt.Print("Deseja Ar condicionado? (s/n): ")
	fmt.Scanln(&ar)
	if ar == "s" {
		precoFinal += 1750
	}

	fmt.Print("Deseja Pintura metálica? (s/n): ")
	fmt.Scanln(&pintura)
	if pintura == "s" {
		precoFinal += 800
	}

	fmt.Print("Deseja Vidro elétrico? (s/n): ")
	fmt.Scanln(&vidro)
	if vidro == "s" {
		precoFinal += 1200
	}

	fmt.Print("Deseja Direção hidráulica? (s/n): ")
	fmt.Scanln(&direcao)
	if direcao == "s" {
		precoFinal += 2000
	}

	fmt.Printf("Preço final do carro: R$ %.2f\n", precoFinal)
}
