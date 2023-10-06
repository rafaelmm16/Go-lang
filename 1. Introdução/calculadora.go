package main

import (
	"fmt"
	"os"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Calculadora em Go")
	fmt.Println("==================")

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	fmt.Print("Digite o operador (+, -, *, /): ")
	fmt.Scanln(&operator)

	result := 0.0
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Erro: Divisão por zero!")
			os.Exit(0)
		}
	default:
		fmt.Println("Operador inválido!")
		os.Exit(0)
	}

	fmt.Printf("Resultado: %f\n", result)
}
