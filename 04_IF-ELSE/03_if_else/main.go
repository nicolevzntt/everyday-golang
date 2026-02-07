package main

import "fmt"

func main() {

	var idade int
	fmt.Println("Qual sua idade? ")
	fmt.Scanf("%d", &idade)

	if idade >= 66 {
		fmt.Println("Olha só ein....")
		fmt.Println("Manere com a bebida!")
	} else if idade >= 18 {
		fmt.Println("Perfeito!")
		fmt.Println("Beba a vontade!")
	} else {
		fmt.Println("Vishh")
		fmt.Println("Melhor voltar pra casa e beber leite")
	}
}
