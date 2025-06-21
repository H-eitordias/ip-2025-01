package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)
	if n >= 9 {
		fmt.Printf("NOTA = %.1f CONCEITO = A\n", n)
	} else if n >= 7.5 {
		fmt.Printf("NOTA = %.1f CONCEITO = B\n", n)
	} else if n >= 6 {
		fmt.Printf("NOTA = %.1f CONCEITO = C\n", n)
	} else {
		fmt.Printf("NOTA = %.1f CONCEITO = D\n", n)
	}
}
