package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Print("Digite os coeficientes A, B e C: ")
	fmt.Scanln(&a, &b, &c)

	if a == 0 {
		fmt.Println("Não é uma equação do 2º grau.")
		return
	}

	delta := b*b - 4*a*c

	if delta < 0 {
		fmt.Println("RAÍZES IMAGINÁRIAS")
	} else if delta == 0 {
		raiz := -b / (2 * a)
		fmt.Printf("RAIZ ÚNICA: %.2f\n", raiz)
	} else {
		r1 := (-b + math.Sqrt(delta)) / (2 * a)
		r2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Printf("RAÍZES DISTINTAS: %.2f e %.2f\n", r1, r2)
	}
}

