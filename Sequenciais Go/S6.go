package main

import "fmt"

func main() {
	var n1 float64
	fmt.Print("Quanto você ganha no seu atual emprego? R: ")
	fmt.Scan(&n1)
	fmt.Print("Se você tivesse um aumento de 15%, iria ganhar: ", n1*1.15)
}
