package main

import "fmt"

//tipo de declaração de string
var name3 string

const idade int = 25

func main() {
	//tipo de declaração de string
	const name string = "rafael"

	//tipo de declaração de string
	name2 := "teste"

	//prints
	fmt.Println("ola", name, idade)
	fmt.Println("ola", name2)
	fmt.Println("ola", name3)
	//T siginifica o tipo do dado
	fmt.Printf("%T %s", name2, name2)
}
