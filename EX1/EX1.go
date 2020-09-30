package main

import "fmt"

func main() {
	var valorpago float64
	var preco float64

	fmt.Print("Informe o valor pago pelo Produto: R$")
	fmt.Scanln(&valorpago)
	fmt.Print("Informe o preço do Produto: R$")
	fmt.Scanln(&preco)

	fmt.Printf("O seu troco é R$ %.2f", valorpago-preco)
}
