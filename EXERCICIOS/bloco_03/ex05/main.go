// Faça um programa que verifique se a pessoa pertence à família "nunes" ou “silva”.

package main

import (
	"fmt"
	"strings"
)

func main () {

var nome string

	fmt.Println("Entre com o seu nome completo:")
	fmt.Scanf("%s", &nome)

	if strings.Contains(nome, "nunes") || strings.Contains(nome, "silva") {

		fmt.Println("Você deve ser da familia Nunes ou Silva")

	} else {
		fmt.Println("Você não é parente de nenhum Nunes ou Silva")
	}

}