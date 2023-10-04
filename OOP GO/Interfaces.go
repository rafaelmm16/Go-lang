package main

import (
	"fmt"
)

type Humano interface {
	Fala(texto string)
	Sente(sentimento string)
}

type Robo struct {
	Nome           string
	DataFabricacao string
}

func (r *Robo) Fala(texto string) {
	fmt.Println(texto)
}

func (p *Pessoa) Sente(sentimento string) {
	fmt.Println("estou sentindo", sentimento)
}

// estrutura
type Pessoa struct {
	Nome  string
	idade int
}

func (p *Pessoa) Fala(texto string) {
	fmt.Println(texto)
}
func CadastroCPF(h Humano) {
	//h.CPF = "12345678978"
	fmt.Println("CPF cadastrado")
}

func main() {

	p := &Pessoa{}
	r := &Robo{}

	p.Fala("HI")
	r.Fala("01")

	CadastroCPF(p)
	//CadastroCPF(r)
}
