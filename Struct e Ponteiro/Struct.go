package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

//recebe estring e retorna valor e uma string
func NewUser(name string) Pessoa {

	p := Pessoa{
		Nome:  name,
		Idade: 13,
	}
	return p
}

func main() {

	//segunda forma de usar a struct
	var pessoas = []Pessoa{
		{"Pedro", 30},
		{"Lipe", 33},
		{"Blue", 36}}

	//quando n precisar de uma informação uso o underline "_"
	for _, v := range pessoas {
		fmt.Printf("Pessoa: %s\t\tIdade: %d\n", v.Nome, v.Idade)
	}

	//uma foram de usar a struct
	name := "rafael"
	pessoa := NewUser(name)
	pessoa2 := NewUser("Jõao")
	pessoa3 := NewUser("Maria")

	fmt.Println(name)
	fmt.Printf("Pessoa: %s\t\tIdade: %d\n", pessoa.Nome, pessoa.Idade)
	fmt.Printf("Pessoa: %s\t\tIdade: %d\n", pessoa2.Nome, pessoa2.Idade)
	fmt.Printf("Pessoa: %s\t\tIdade: %d\n", pessoa3.Nome, pessoa3.Idade)
}
