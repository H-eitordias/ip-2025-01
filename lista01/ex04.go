package main

import "fmt"

func main() {
	var s, k float64
	fmt.Scan(&s, &k)
	v := s * 0.7 / 100
	t := v * k
	d := t * 0.9
	fmt.Printf("Custo por kW: R$ %.2f\n", v)
	fmt.Printf("Custo do consumo: R$%.2f\n", t)
	fmt.Printf("Custo com desconto: R$ %.2f\n", d)
}
