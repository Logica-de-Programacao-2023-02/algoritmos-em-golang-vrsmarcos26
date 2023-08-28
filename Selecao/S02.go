package main

import (
	"fmt"
)

func main() {
	var n1, n2, n3 int
	fmt.Print("Me fale um número: ")
	fmt.Scan(&n1)
	fmt.Print("Me fale um número: ")
	fmt.Scan(&n2)
	fmt.Print("Me fale um número: ")
	fmt.Scan(&n3)

	if n1 < n2 && n1 < n3 {
		fmt.Println("O menor número entre eles é: ", n1)
	}
	if n2 < n1 && n2 < n3 {
		fmt.Println("O menor número entre eles é: ", n2)
	}
	if n3 < n1 && n3 < n2 {
		fmt.Println("O menor número entre eles é: ", n3)
	}
}
