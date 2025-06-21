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

	switch opcao {
	case 1: // Cone
		volume := (math.Pi * math.Pow(raio, 2) * altura) / 3
		area := math.Pi * raio * math.Sqrt(math.Pow(raio, 2)+math.Pow(altura, 2))
		fmt.Printf("Volume: %.2f\nÁrea: %.2f\n", volume, area)
	case 2: // Cilindro
		volume := math.Pi * math.Pow(raio, 2) * altura
		area := 2 * math.Pi * raio * altura
		fmt.Printf("Volume: %.2f\nÁrea: %.2f\n", volume, area)
	case 3: // Esfera
		volume := (4.0 / 3.0) * math.Pi * math.Pow(raio, 3)
		area := 4 * math.Pi * math.Pow(raio, 2)
		fmt.Printf("Volume: %.2f\nÁrea: %.2f\n", volume, area)
	default:
		fmt.Println("Opção inválida.")
	}
}

