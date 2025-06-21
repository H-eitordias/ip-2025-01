package main

import (
	"fmt"
)

func main() {
	var matricula, horas int
	const salarioMinimo = 788.0
	const valorHoraExtra = 10.0

	fmt.Print("Matrícula do funcionário: ")
	fmt.Scanln(&matricula)
	fmt.Print("Horas extras: ")
	fmt.Scanln(&horas)

	salarioBruto := 3*salarioMinimo + float64(horas)*valorHoraExtra
	inss := 0.0
	if salarioBruto > 1500 {
		inss = salarioBruto * 0.12
	}
	ir := 0.0
	if salarioBruto > 2000 {
		ir = salarioBruto * 0.20
	}
	salarioLiquido := salarioBruto - inss - ir

	fmt.Printf("Bruto: R$ %.2f | Líquido: R$ %.2f\n", salarioBruto, salarioLiquido)
}
