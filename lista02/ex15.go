package main

import (
	"fmt"
)

func main() {
	var dia, mes, ano int
	fmt.Print("Digite o dia: ")
	fmt.Scanln(&dia)
	fmt.Print("Digite o mês (1 a 12): ")
	fmt.Scanln(&mes)
	fmt.Print("Digite o ano: ")
	fmt.Scanln(&ano)

	meses := []string{
		"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho",
		"Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro",
	}

	if mes < 1 || mes > 12 {
		fmt.Println("Mês inválido.")
		return
	}

	fmt.Printf("%d de %s de %d\n", dia, meses[mes-1], ano)
}
