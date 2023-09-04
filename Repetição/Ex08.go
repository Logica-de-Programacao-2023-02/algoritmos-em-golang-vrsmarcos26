package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&n)
	if n <= 0 {
		fmt.Println("Por favor, insira um número inteiro positivo.")
	} else {
		fmt.Printf("Os divisores de %d são:\n", n)
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				fmt.Println(i)
			}
		}
	}
}
