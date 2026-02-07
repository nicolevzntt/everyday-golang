package main

import "fmt"

func main() {

	nome := "Paramore"

	for _, v := range nome {

		letra := string(v)
		if letra == "m" {
			break
		}

		if letra == " " {
			continue
		}

		fmt.Println(letra)
	}
}

//fmt.Println("teste")
