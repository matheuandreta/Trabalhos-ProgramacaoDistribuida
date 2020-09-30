package main

import (
	"fmt"
)

func main() {

	var A bool
	var B float64 = 10.50
	var C int = 1
	var D string = "Ponteiros"

	fmt.Println("Valor da string: ", D)
	fmt.Println("Endereço de memória string: ", &D)

	fmt.Println("Valor do float64: ", B)
	fmt.Println("Endereço de memória float64: ", &B)

	fmt.Println("Valor do int: ", C)
	fmt.Println("Endereço de memória int: ", &C)

	fmt.Println("Valor do bool: ", A)
	fmt.Println("Endereço de memória bool: ", &A)

}
