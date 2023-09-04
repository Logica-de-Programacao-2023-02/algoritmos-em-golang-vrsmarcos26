package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Me diga um número: ")
	fmt.Scan(&n1)
	fmt.Print("Me diga mais um número: ")
	fmt.Scan(&n2)

	if n1 >= 0 && n2 >= 0 {
		fmt.Println("Resultado da multiplicação", n1*n2)
	}
	if n1 < 0 || n2 < 0 {
		fmt.Println("O resultado da soma será: ", n1+n2)
	}
}
