package main

import "fmt"

//recebe estring e retorna valor e uma string
func NewUser(name string) (string, int) {
	fmt.Println("Recebi: ", name)
	name = "Valor dentro da func"
	fmt.Println(name)
	return "Rafael", 25
}

func main() {
	fmt.Println("hi")

	name := "rafael"
	pessoa, idade := NewUser(name)
	fmt.Println(name)
	fmt.Printf("Pessoa %s\n", pessoa)
	fmt.Printf("Idade %d\n", idade)
}
