package main

import "fmt"

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)

det := a*d - b*c
	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f\n", det)
}
