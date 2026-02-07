// faça um programa que exiba o dobro de um numero inseirdo pelo usuario
// receba um numero inteiro e calcule sua razia quadrada e exiba o resultado

package main

import (
	"fmt"
)

func main() {

	var numero float64

	fmt.Print("Entre com um numero:")
	fmt.Scanf("%f", &numero)

	res := 2 * numero
	fmt.Printf("O dobro de %.2f é %.2f\n", numero, res)

}
