// Faça um programa que receba 4 alturas usando um laço de repetição e realize a soma dessas alturas.

package main

import "fmt"

func main() {

	var altura, soma float64 // iniciadas em 0

	for i := 1; i <= 4; i++ {

		fmt.Printf("Entre com %da altura: ", i)
		fmt.Scanln(&altura)

		soma += altura
	}

	fmt.Printf("A soma de todas as alturas é: %.2f", soma)
}
