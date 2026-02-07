//Faça um programa que vende uma garrafa de água:
//Se o cliente escolher água mineral natural, será cobrado R$1,50
//Se o cliente escolher água mineral com gás, será cobrado R$2,50

package main

import "fmt"

func main() {

	var opcao int

	fmt.Println("Qual água você deseja comprar? Mineral (1) ou Com Gás (2)")
	fmt.Scanf("%d", &opcao) // Leia um inteiro e guarde no endereço de opcao

	txtTemplate := "o valor ficou: R$%.2f\n" // a variavel guarda um modelo de texto

	switch opcao {
	case 1:
		fmt.Printf(txtTemplate, 1.5)

	case 2:
		fmt.Printf(txtTemplate, 2.5)

	default:
		fmt.Println("Digite um valor válido!")
	}

	/* 	var opcao int

	   	fmt.Println("Qual água você deseja comprar? Mineral (1) ou Com Gás (2)")
	   	fmt.Scanf("%d", &opcao)

	   	var valor float64
	   	if opcao == 1 {
	   		valor = 1.5
	   	} else if opcao == 2 {
	   		valor = 2.5
	   	}

	   	if valor != 0 {
	   		fmt.Printf("O valor ficou: R$%.2f\n", valor)
	   	} else {
	   		fmt.Println("Digite um valor válido!")
	   	} */

}
