package main

import (
	"fmt"
)

// estrutura
type Site struct {
	Criador string
	URL     string
}

// metodo
func (s *Site) DefineCriador(name string) {
	s.Criador = name
	fmt.Println("Dado salvo")
}
func (s *Site) DefineUrl(URL string) {
	s.URL = URL
	fmt.Println("Dado salvo")
}

func main() {

	site := Site{}

	site.DefineCriador("Rafael")
	site.DefineUrl("https://github.com/rafaelmm16/Go-lang")
	fmt.Printf("Site: %s\nCriador: %s\n", site.URL, site.Criador)

	site.DefineCriador("Rafael2")
	site.DefineUrl("https://github.com/rafaelmm16/Go-lang")
	fmt.Printf("Site: %s\nCriador: %s\n", site.URL, site.Criador)
}
