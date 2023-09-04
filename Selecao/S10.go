package main

import "fmt"

func main() {
	var n1 float64
	fmt.Print("Qual sua idade? R: ")
	fmt.Scan(&n1)
	if n1 <= 9 {
		fmt.Print("Sua classificação é MIRIM")
	} else if n1 <= 13 {
		fmt.Print("Sua classificação é INFANTIL")
	} else if n1 <= 7 {
		fmt.Print("Sua classificação é JUVENIL")
	} else if n1 >= 18 {
		fmt.Print("Sua classificação é ADULTO")
	}
}
