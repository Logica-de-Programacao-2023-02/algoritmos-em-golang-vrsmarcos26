package main

import "fmt"

func main() {
	var M, N1 int
	N1 = -1
	for N1 != 0 {

		fmt.Print("Digite um número: ")
		fmt.Scan(&N1)
		if N1 > M {
			M = N1
		}
	}
	fmt.Print("A média é ", M)

}
