package main

import "fmt"

func main() {
	var h int
	fmt.Scan(&h)
	v := (h/3)*10 + (h%3)*5
	fmt.Printf("O VALOR A PAGAR E = %.2f\n", float64(v))
}
