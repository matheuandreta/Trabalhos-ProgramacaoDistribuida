package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64

	fmt.Print("Digite o valor de A:")
	fmt.Scanln(&a)
	fmt.Print("Digite o valor de B:")
	fmt.Scanln(&b)
	fmt.Print("Digite o valor de C:")
	fmt.Scanln(&c)
	fmt.Print("Digite o valor de D:")
	fmt.Scanln(&d)
	fmt.Println("----------------------")
	menor := math.Min(float64(a), float64(b))
	menor = math.Min(float64(menor), float64(c))
	menor = math.Min(float64(menor), float64(d))

	fmt.Println("O menor valor entre A, B, C e D é: ", menor)

}
