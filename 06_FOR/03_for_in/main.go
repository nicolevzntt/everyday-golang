package main

import "fmt"

func main() {

	nome := "GloboPlay"

	// range está acessando os indices dos bytes da string
	for i := range nome {
		fmt.Println(i, nome[i], string(nome[i])) // indice; byte na posição que corresponde ao caracter;
		// string(nome[i]):Pega esse byte e transforme numa string de 1 caractere
	}
}
