package main

import "fmt"

func main() {

	// "1- banana"
	// "2- maça"
	// "3- pera"
	// "4- abacaxi"

	var opcao int

	fmt.Println("Escolha uma opção: ")
	fmt.Scanf("%d", &opcao)

	switch opcao {

	case 1:
		fmt.Println("Vc escolheu banana!")

	case 2:
		fmt.Println("Vc escolheu maça!")

	case 3:
		fmt.Println("Vc escolheu pera!")

	case 4:
		fmt.Println("Vc escolheu abacaxi!")

	default:
		fmt.Println("Opção inválida!")
	}
}
