// faça um programa que receba um numero inteiro
// e calcule sua razia quadrada e exiba o resultado

package main

import (
	"fmt"
	"math"
)

func main() {

	var numero int
	fmt.Println("Digite um numero: ")
	fmt.Scanf("%d", &numero)

	resultado := math.Sqrt(float64(numero)) //add a biblioteca math porem ela exige que o numero seja float e
	//aqui convertemos a variavel numero que esta em int para float64, e obtemos o resultado da raiz
	// resultado recebe o resultado de math.Sqrt(float64(numero))

	fmt.Printf("A raiz quadrada de %d é %.2f\n", numero, resultado)

}
