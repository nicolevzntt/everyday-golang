package main

import "fmt"

func main() {

	// // for estilo while (menos legível)
	// i := 1
	// for i <= 10 {
	// 	fmt.Println("For all I know")
	// 	fmt.Println("The best is over and the worst is yet to come")
	// 	i++
	//}

	for i := 26; i <= 1000; i++ {
		if i%25 == 0 {
			fmt.Println("Continua a busca", i)
			fmt.Println("Encontrei um múltiplo de 25")
			break
		}
	}
}
