package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		var total int
		var p, g, a, c float64
		fmt.Scan(&total, &p, &g, &a, &c)
		r := float64(total) * (p*1 + g*5 + a*10 + c*20) / 100
		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i, r)
	}
}
