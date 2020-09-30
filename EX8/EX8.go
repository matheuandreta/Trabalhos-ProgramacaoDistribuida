package main

import "fmt"

func main() {
	var ladoA, ladoB, ladoC int

	fmt.Print("Digite o tamanho do lado A do triangulo: ")
	fmt.Scanln(&ladoA)
	fmt.Print("Digite o tamanho do lado B do triangulo: ")
	fmt.Scanln(&ladoB)
	fmt.Print("Digite o tamanho do lado C do triangulo: ")
	fmt.Scanln(&ladoC)

	if ladoA == 0 || ladoB == 0 || ladoC == 0 {
		fmt.Println("Com essas medidas não é possivel fomar um triângulo")
	} else {
		if ladoA > (ladoB+ladoC) || ladoB > (ladoA+ladoC) || ladoC > (ladoA+ladoB) {
			fmt.Println("Com essas medidas não é possivel fomar um triângulo")
		} else {
			if ladoA == ladoB && ladoA == ladoC {
				fmt.Println("Esse triângulo é equilátero")

			} else {

				if ladoA == ladoB || ladoB == ladoC || ladoA == ladoC {
					fmt.Println("Esse triângulo é isósceles")

				}

			}

			if ladoA != ladoB && ladoB != ladoC && ladoA != ladoC {
				fmt.Println("Esse triângulo é escaleno")

			}
		}

	}

}
