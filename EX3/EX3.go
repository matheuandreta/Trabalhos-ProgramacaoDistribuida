package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Print("Digite o valor de A:")
	fmt.Scanln(&a)
	fmt.Print("Digite o valor de B:")
	fmt.Scanln(&b)
	fmt.Println("----------------------")
	fmt.Println("O maior valor entre a e b é: ", math.Max(float64(a), float64(b)))

}
