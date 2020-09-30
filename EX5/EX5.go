package main

import (
	"fmt"
)

func main() {
	var lista1 [10]int
	var lista2 [10]int

	for i := 0; i < 10; i++ {
		fmt.Print("Digite um numero:")
		fmt.Scanln(&lista1[i])

	}
	var j = 9
	for i := 0; i < 10; i++ {

		lista2[j] = lista1[i]
		j = j - 1

	}

	fmt.Println("Numeros digitados:", lista1)
	fmt.Println("Numeros invertidos", lista2)

}
