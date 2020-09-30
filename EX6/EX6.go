package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	var lista1 [50]int
	var numero int

	for i := 0; i < 50; i++ {
		time.Sleep(1 * time.Nanosecond)
		rand.Seed(time.Now().UnixNano())
		lista1[i] = randomInt(1, 200)
	}

	for i := 0; i < 50; i++ {
		if lista1[i] >= 10 && lista1[i] <= 100 {
			numero = numero + 1
		} else {

		}

	}

	fmt.Println(lista1)
	fmt.Println("Quantidade de numeros entre o 10 e o 100: ", numero)

}
