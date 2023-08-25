package main

import "fmt"

func main() {
	var n1 float64
	fmt.Print("Quantos quilos você está pesando atualmente? R:")
	fmt.Scan(&n1)
	fmt.Print("Seu peso em libras atualmente é:", n1*2.20462)
}
