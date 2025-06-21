package main

import (
	"fmt"
)

func main() {
	var tipo string
	var consumo float64
	var valor float64

	fmt.Print("Tipo de cliente (residencial, comercial, industrial): ")
	fmt.Scanln(&tipo)
	fmt.Print("Consumo em m³: ")
	fmt.Scanln(&consumo)

	switch tipo {
	case "residencial":
		valor = 5 + 0.05*consumo
	case "comercial":
		if consumo <= 80 {
			valor = 500
		} else {
			valor = 500 + 0.25*(consumo-80)
		}
	case "industrial":
		if consumo <= 100 {
			valor = 800
		} else {
			valor = 800 + 0.04*(consumo-100)
		}
	default:
		fmt.Println("Tipo inválido.")
		return
	}

	fmt.Printf("Valor a pagar: R$ %.2f\n", valor)
}
