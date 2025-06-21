 package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Print("Digite três números inteiros diferentes: ")
	fmt.Scanln(&a, &b, &c)

	var menor, inter, maior int

	// Comparações para definir a ordem
	if a < b && a < c {
		menor = a
		if b < c {
			inter = b
			maior = c
		} else {
			inter = c
			maior = b
		}
	} else if b < a && b < c {
		menor = b
		if a < c {
			inter = a
			maior = c
		} else {
			inter = c
			maior = a
		}
	} else {
		menor = c
		if a < b {
			inter = a
			maior = b
		} else {
			inter = b
			maior = a
		}
	}

	fmt.Printf("Ordem crescente: %d, %d, %d\n", menor, inter, maior)
}
