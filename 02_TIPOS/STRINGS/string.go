package main

import "fmt"

func main() {

	fmt.Println("O que voce quiser!!!")

	fmt.Println(`Aqui a gente pode escrever múltiplas linhas de código!`)

	fmt.Println("nicole tem", len("nicole"), "letras") // len vem de length (comprimento).Ele serve para dizer quantos elementos algo tem.

	fmt.Println("Nicole"[0]) // isso é um byte (int8)
	
	fmt.Println(string("Nicole"[0]))
	fmt.Println(string("Nicole"[1]))
	fmt.Println(string("Nicole"[2]))
	fmt.Println(string("Nicole"[3]))

}
