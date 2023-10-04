package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	resp, err := http.Get("https://651ded2844e393af2d5a6c82.mockapi.io/api/v1/users")
	if err != nil {
		fmt.Println("Aconteceu um erro:", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	fmt.Println(string(body))
}
