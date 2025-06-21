package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Print("Digite dois números inteiros (A e B): ")
	fmt.Scanln(&a, &b)

	if b == 0 {
		fmt.Println("Divisão por zero não é permitida.")
	} else if a%b == 0 {
		fmt.Println("A é divisível por B.")
	} else {
		fmt.Println("A não é divisível por B.")
	}
}
