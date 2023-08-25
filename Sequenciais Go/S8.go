package main

import "fmt"

func main() {
	var n1, n2 float64
	fmt.Println("Vamos calcular seu salário.")
	fmt.Print("Primeiro, conte-me quantos dias você trabalhou:")
	fmt.Scan(&n1)
	fmt.Print("Agora me conte o valor da sua diária:")
	fmt.Scan(&n2)
	fmt.Print("No final do mês seu salário será de:", n1*n2)
}
