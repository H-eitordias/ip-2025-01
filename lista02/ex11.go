package main

import (
	"fmt"
)

func main() {
	var x float64
	fmt.Print("Digite o valor de x: ")
	fmt.Scanln(&x)

	if x == 2 {
		fmt.Println("Divisão por zero! f(x) é indefinido para x = 2.")
	} else {
		fx := 8 / (2 - x)
		fmt.Printf("f(x) = %.2f\n", fx)
	}
}
