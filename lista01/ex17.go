package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x%2 != 0 {
		fmt.Println("O PRIMEIRO NUMERO NAO E PAR")
	} else {
		fmt.Print(x)
		if y > 1 {
			fmt.Print(" ", x+2)
		}
		if y > 2 {
			fmt.Print(" ", x+4)
		}
		if y > 3 {
			fmt.Print(" ", x+6)
		}
		if y > 4 {
			fmt.Print(" ", x+8)
		}
		fmt.Println()
	}
}
