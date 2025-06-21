package main

import (
	"fmt"
)

func main() {
	var valorCompra, valorVenda float64
	fmt.Print("Digite o valor da compra: R$ ")
	fmt.Scanln(&valorCompra)

	if valorCompra < 0 {
		fmt.Println("Valor inválido.")
	} else if valorCompra < 10 {
		valorVenda = valorCompra * 1.70
	} else if valorCompra < 30 {
		valorVenda = valorCompra * 1.50
	} else if valorCompra < 50 {
		valorVenda = valorCompra * 1.40
	} else {
		valorVenda = valorCompra * 1.30
	}

	fmt.Printf("Valor da venda: R$ %.2f\n", valorVenda)
}
