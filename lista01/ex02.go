package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a > 9 || b > 9 || c > 9 {
		fmt.Println("DIGITO INVALIDO")
		return
	}
	n := a*100 + b*10 + c
	fmt.Printf("%d, %d\n", n, n*n)
}
