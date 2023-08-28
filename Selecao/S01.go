package main

import "fmt"

func main() {
	var N1, N2 int
	fmt.Print("Me informe um número: ")
	fmt.Scan(&N1)
	fmt.Print("Me informe um número: ")
	fmt.Scan(&N2)

	if N1 > N2 {
		fmt.Println("O maior número entre os dois é : ", N1)
	} else {
		fmt.Println("O maior número entre os dois é : ", N2)
	}
}
