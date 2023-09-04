package main

import "fmt"

func main() {
	var n1 float64
	fmt.Print("Qual sua idade? R: ")
	fmt.Scan(&n1)
	if n1 <= 9 {
		fmt.Print("Sua classificação é MIRIM")
	}
	if n1 == 10 || n1 == 11 || n1 == 12 || n1 == 13 {
		fmt.Print("Sua classificação é INFANTIL")
	}
	if n1 == 14 || n1 == 15 || n1 == 16 || n1 == 17 {
		fmt.Print("Sua classificação é JUVENIL")
	}
	if n1 >= 18 {
		fmt.Print("Sua classificação é ADULTO")
	}
}
