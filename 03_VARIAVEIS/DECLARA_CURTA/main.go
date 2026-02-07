package main

import "fmt"

func main() {

	x := 45

	nome := "Nicole"

	y := 0.45

	fmt.Println(x, nome, y)

	x = 90
	fmt.Println(x) // 👉 x já existe, o tipo já é conhecido (int), então você só muda o valor

	fmt.Printf("x = %T; nome = %T ; y = %T \n", x, nome, y)

}
