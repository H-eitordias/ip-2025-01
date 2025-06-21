package main

import "fmt"

func main() {
	var c int
	var m float64
	var t byte
	fmt.Scan(&c, &m, &t)
	v := 0.0

	if t == 'R' {
		v = 5 + 0.05*m
	} else if t == 'C' {
		v = 500
		if m > 80 {
			v += (m - 80) * 0.25
		}
	} else if t == 'I' {
		v = 800
		if m > 100 {
			v += (m - 100) * 0.04
		}
	}

	fmt.Printf("CONTA = %d\n", c)
	fmt.Printf("VALOR DA CONTA = %.2f\n", v)
}
