package main

import (
	"fmt"
	"math"
)

func main() {
	var numero float64
	fmt.Print("Digite um número: ")
	fmt.Scanln(&numero)

	if numero >= 0 {
		fmt.Println("Raiz quadrada:", math.Sqrt(numero))
	} else {
		fmt.Println("Quadrado:", numero*numero)
	}
}
