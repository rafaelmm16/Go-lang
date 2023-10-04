package main

import "fmt"

func main() {
	fmt.Println("hi")

	name := "rafaell"

	if name == "rafael" {
		fmt.Println("É vdd", name)
	} else if name == "rafaell" {
		fmt.Println("É vdd 2", name)
	} else {
		fmt.Println("É falso")
	}
}
