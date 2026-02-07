// Faça um programa que a pessoa entra com 3 notas e
// verificamos se a média delas é suficiente para passar na prova. Nota de corte 6.

package main

import "fmt"

func main() {

	var nota, soma float64

	fmt.Println("Digite a nota da prova 1:")
	fmt.Scan(&nota)
	soma += nota

	fmt.Println("Digite a nota da prova 2:")
	fmt.Scan(&nota)
	soma = soma + nota

	fmt.Println("Digite a nota da prova 3:")
	fmt.Scan(&nota)
	soma += nota // pega o valor que já está em soma e adicione nota a ele.

	media := soma / 3.0

	if media >= 6 {
		fmt.Printf("Voce passouu!! Media = %2.f\n", media)
	} else {
		fmt.Printf("Vai ter que estudar mais! Media = %2.f\n", media)
	}

}
