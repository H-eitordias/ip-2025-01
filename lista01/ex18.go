package main

import "fmt"

func main() {
	var a1, r, n int
	fmt.Scan(&a1, &r, &n)
	soma := a1 + (a1+r)
	if n > 2 {
		soma += a1 + 2*r
	}
	if n > 3 {
		soma += a1 + 3*r
	}
	if n > 4 {
		soma += a1 + 4*r
	}
	fmt.Println(soma)
}
