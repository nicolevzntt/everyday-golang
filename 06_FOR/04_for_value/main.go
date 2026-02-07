package main

import "fmt"

// func main() {

// 	nome := "GloboPlay"

// 	// v é o valor numérico do caractere unicode
// 	// for i, v := range nome {
// 	// 	fmt.Println(i, v, string(v))
// 	// }

func main() {

	nome := "GloboPlay"

	// _ signifcia que o índice existe, mas descarta e guarda só o valor do caractere em v
	// ou seja sei que existe, mas não quero que mostre
	for _, v := range nome {
		fmt.Println(string(v)) // (string(v) pega o número Unicode e transforme no caractere correspondente
	}
}

// or:
// func main() {

// 	nome := "GloboPlay"

// 	for _, letraByte := range nome {
// 		letra := string(letraByte)
// 		fmt.Println(letra)
// 	}
