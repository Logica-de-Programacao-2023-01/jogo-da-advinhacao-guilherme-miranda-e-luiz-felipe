package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	num := rand.Intn(100) + 1

	fmt.Println("Tente acertar o número entre 1 e 100")

	var tries int
	for {
		var n int
		fmt.Print("Digite um número:")
		fmt.Scan(&n)

		tries++

		if n == num {
			fmt.Printf("Você acertou! Utilizou %d tentativas.\n", tries)
			break
		} else if n < num {
			fmt.Println("Digite um número maior.")
		} else {
			fmt.Println("Digite um número menor.")
		}
	}
}
