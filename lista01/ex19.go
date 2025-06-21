package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n < 2 {
		fmt.Println("Numero invalido!")
		return
	}
	s := 1.0
	if n >= 2 {
		s += 1.0 / 2
	}
	if n >= 3 {
		s += 1.0 / 3
	}
	if n >= 4 {
		s += 1.0 / 4
	}
	if n >= 5 {
		s += 1.0 / 5
	}
	fmt.Printf("%.6f\n", s)
}
