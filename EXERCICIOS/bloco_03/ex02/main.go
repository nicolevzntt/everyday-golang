//Altere o programa anterior para considerar a quantidade de garrafas de água

package main

import "fmt"

func main() {

	var opcao, quantidade int
	fmt.Println("Qual água você deseja comprar? Mineral (1) ou Com Gás (2)")
	fmt.Scanf("%d", &opcao)

	fmt.Println("Quantas garrafas vc quer? ")
	fmt.Scanf("%d", &quantidade)

	txtTemplate := "O valor ficou: R$%.2f"
	// ou txtTemplate := "o valor ficou: R$%.2f x %d = R$%.2f\n"

	switch opcao {
	case 1:
		fmt.Printf(txtTemplate, 1.5*float64(quantidade))
		//fmt.Printf(txtTemplate, 1.5, quantidade, 1.5*float64(quantidade))

	case 2:
		fmt.Printf(txtTemplate, 2.5*float64(quantidade))

	default:
		fmt.Println("Digite um valor válido!")
	}

}
