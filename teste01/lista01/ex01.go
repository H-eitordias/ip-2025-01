package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	m := (a + b + c) / 3
	fmt.Printf("MEDIA = %.2f\n", m)
	if m >= 6 {
		fmt.Println("APROVADO")
	} else {
		fmt.Println("REPROVADO")
	}
}
