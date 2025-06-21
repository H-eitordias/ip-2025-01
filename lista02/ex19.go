package main

import (
	"fmt"
	"math"
)

func main() {
	var opcao int
	var raio, altura float64

	fmt.Println("1 - Cone\n2 - Cilindro\n3 - Esfera")
	fmt.Print("Escolha a figura: ")
	fmt.Scanln(&opcao)

	fmt.Print("Digite o raio: ")
	fmt.Scanln(&raio)

	if opcao == 1 || opcao == 2 {
		fmt.Print("Digite a altura: ")
		fmt.Scanln(&altura)
	}

	if opcao == 1 { 
		v := math.Pi * raio * raio * altura / 3
		a := math.Pi * raio * math.Sqrt(raio*raio+altura*altura)
		fmt.Printf("Cone - Volume: %.2f, Área: %.2f\n", v, a)
	} else if opcao == 2 {
		v := math.Pi * raio * raio * altura
		a := 2 * math.Pi * raio * altura
		fmt.Printf("Cilindro - Volume: %.2f, Área: %.2f\n", v, a)
	} else if opcao == 3 { 
		v := 4 * math.Pi * math.Pow(raio, 3) / 3
		a := 4 * math.Pi * raio * raio
		fmt.Printf("Esfera - Volume: %.2f, Área: %.2f\n", v, a)
	} else {
		fmt.Println("Opção inválida.")
	}
}
