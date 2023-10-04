package main

import "fmt"

// DizBomdia é uma função que imprime uma mensagem
func DizBomdia() {
	fmt.Println("Bom dia")
}

func main() {
	indice := 0
	for true {
		fmt.Println("Indice atual:", indice)
		indice = indice + 1
		if indice > 10 {
			fmt.Println("Loop passou de 10")

			// Quando o loop passar de 10, diz Bom dia
			DizBomdia()

			// sai do loop
			break
		}
	}
}
