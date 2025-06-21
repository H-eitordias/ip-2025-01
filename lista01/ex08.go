package main

import "fmt"

func main() {
	var r, h float64
	fmt.Scan(&r, &h)
	pi := 3.14159
	area := 2*pi*r*r + 2*pi*r*h
	custo := area * 100
	fmt.Printf("O VALOR DO CUSTO E = %.2f\n", custo)
}
