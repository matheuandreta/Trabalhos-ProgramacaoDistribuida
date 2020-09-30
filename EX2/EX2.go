package main

import "fmt"

func main() {
	var n1, resto int

	fmt.Print("Digite um numero:")
	fmt.Scanln(&n1)

	resto = n1 % 2

	if resto == 0 {
		fmt.Println("Este numero não é impar")

	} else {
		fmt.Println("Esse numero é impar")
	}

}
