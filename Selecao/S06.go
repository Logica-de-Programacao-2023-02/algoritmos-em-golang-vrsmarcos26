package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Me diga um número: ")
	fmt.Scan(&n1)
	fmt.Print("Me diga mais um número: ")
	fmt.Scan(&n2)

	if n1 >= 0 && n2 >= 0 {
		resultado := n1 * n2
		fmt.Println("Resultado da multiplicação", resultado)
	}
	if n1 < 0 || n2 < 0 {
		resultado := n1 + n2
		fmt.Println("O resultado da soma será: ", resultado)
	}
}
