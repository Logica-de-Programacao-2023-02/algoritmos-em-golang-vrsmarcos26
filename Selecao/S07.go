package main

import "fmt"

func main() {
	var n1 float64
	fmt.Print("Me conte o seu salário: ")
	fmt.Scan(&n1)
	if n1 < 1000 {
		fmt.Print("Seu salário terá um aumento de 10% e será: ", n1*1.1)
	}
	if n1 >= 1000 {
		fmt.Print("Seu salário terá um aumento de 5% e será: ", n1*1.05)
	}
}
