//Faça um programa que conte quantas vezes a letra “a” aparece em uma palavra

package main

import (
	"fmt"
	"strings"
)

func main() {

	var palavra string
	fmt.Println("Entre com uma palavra: ")
	fmt.Scanf("%s", &palavra)

	palavra = strings.ToLower(palavra)

	contador := 0
	for _, letraByte := range palavra {

		if string(letraByte) == "a" { // letraByte não é uma string, por isso, a converção
			contador++
		}
	}

	fmt.Printf("A palavra '%s' tem %d As\n", palavra, contador)

}
