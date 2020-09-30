package main

import "fmt"

func main()  {

	var produto string
	produto = "panetone"

	var P *string
	P = &produto

	fmt.Println("O valor de produto é", produto)
	fmt.Println("O endereçe de produto é", &produto)
	fmt.Println("O endereço de P é", &P)
	fmt.Println("O valor de P é", P)
	fmt.Println("O valor apontado por P é", &produto)


	
}