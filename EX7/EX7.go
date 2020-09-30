package main

import "fmt"

func main() {
	var lista1 [50]int
	var maior, par int

	for i := 0; i < 50; i++ {
		fmt.Print("Digite um numero para N", i, ":")
		fmt.Scanln(&lista1[i])
	}

	for i := 0; i < 50; i++ {

		par = lista1[i] % 2

		if par == 0 && maior < lista1[i] {
			maior = lista1[i]

		} else {

		}

	}

	fmt.Println(lista1)
	fmt.Println("Maior numero par digitado: ", maior)

}
