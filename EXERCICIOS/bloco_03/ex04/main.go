// Faça um programa que verifique se a pessoa pertence à família "nunes”.

package main

import (
	"fmt"
	"strings"
)

func main() {

	// var resposta string

	// fmt.Println("Pertence a familia nunes?")
	// fmt.Scan(&resposta)

	// if resposta == "sim" {
	// 	fmt.Println("Pertence a família nunes!!!")
	// } else {
	// 	fmt.Println("Não pertence a familia nunes")
	// }

	var nome string

	fmt.Println("Entre com o seu nome completo:")
	fmt.Scanf("%s", &nome)

	//serve pra verificar se um texto contém outro texto. passamos uma string e uma sub-string e retorna
	//um bool
	if strings.Contains(nome, "nunes") {

		fmt.Println("Você deve ser da familia Nunes")

	} else {
		fmt.Println("Você não é parente de nenhum Nunes")
	}
}
