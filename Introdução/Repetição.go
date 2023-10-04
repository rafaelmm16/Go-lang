package main

import "fmt"

func main() {
	fmt.Println("hi")

	//definição de variaveis
	var nomes = []string{"Carlos", "Maria", "Clara"}
	var idade = []int{14, 22, 36}

	//laõ for
	for i := 0; i < 10; i++ {
		fmt.Println("Número: ", i)
	}

	for i, v := range nomes {
		fmt.Printf("Ola %s, posição: %d e idade %d \n", v, i, idade[i])
	}
}
