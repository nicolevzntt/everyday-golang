// Faça um programa que receba uma quantidade indefinida de valores correspondentes a “saldo em conta”,
// mas quando o usuário apertar “enter” sem digitar valor algum, o programa para de receber valores,
// e exibe a soma de todos os valores digitados anteriormente.

package main

import (
	"fmt"
	"strconv" // pacote strconv é usado para converter strings em números e vice-versa
)

func main() {

	var soma float64

	for {
		var entrada string
		fmt.Print("Entre com um valor: ")
		fmt.Scanln(&entrada)

		if entrada == "" {
			break
		}

		valor, erro := strconv.ParseFloat(entrada, 64) //strconv.ParseFloat(entrada, 64) → converte a string digitada em um número decimal (float64). Retorna dois valores:
		if erro != nil {                               // Verifica se houve erro na conversão.
			fmt.Println("Entre com um valor válido!")
			continue //pula para a próxima iteração do loop, pedindo a entrada novamente.
		}
		soma += valor
	}
	fmt.Printf("Saldo total em conta: R$%.2f", soma)
}
