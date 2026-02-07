//Faça o programa de uma sorveteria, onde o usuário pode escolher:
//Tipo de sorvete: casquinha (R$1,00), cascão (R$2,50), cestinha (R$4,00)
//Sabor do sorvete: morango, creme, chocolate
//Cobertura: Caramelo (R$1,50), morango (R$1,50), chocolate (R$1,50), sem cobertura (R$0,00).
//Apresente o valor a ser pago

package main

import "fmt"

func main() {

	var sorvete, sabor, cobertura string

	fmt.Print("Escolha o tipo de sorvete [casquinha/cascao/cestinha]\n> ")
	fmt.Scan(&sorvete)

	fmt.Print("Escolha o sabor do sorvete [morango/creme/chocolate]\n> ")
	fmt.Scan(&sabor)

	fmt.Print("Escolha a cobertura [caramelo/morango/chocolate/sem]\n> ")
	fmt.Scan(&cobertura)

	var valorSorvete, valorCobertura float64

	switch sorvete {
	case "casquinha":
		valorSorvete = 1.00
	case "cascao":
		valorSorvete = 2.50
	case "cestinha":
		valorSorvete = 4.00
	default:
		fmt.Println("Opção inválida para tipos de sorvete!!")
	}

	switch cobertura {
	case "caramelo", "morango", "chocolate":
		valorCobertura = 1.50
	case "sem":
		valorCobertura = 0
	default:
		fmt.Println("Opção inválida para cobertura")
	}

	valorTotal := valorSorvete + valorCobertura
	fmt.Printf("Total: R$ %.2f\n", valorTotal)
}
