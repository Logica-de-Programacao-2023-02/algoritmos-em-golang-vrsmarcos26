package main

import "fmt"

func main() {
	var n1 int

	fmt.Print("Me mostre um número e mostrarei se ele é multiplo de 3 e de 5 ao mesmo tempo: ")
	fmt.Scan(&n1)

	if n1%3 == 0 && n1%5 == 0 {
		fmt.Print("Seu número é muliplo de 3 e 5")
	} else {
		fmt.Print("Seu não número é muliplo de 3 e 5")
	}
}
