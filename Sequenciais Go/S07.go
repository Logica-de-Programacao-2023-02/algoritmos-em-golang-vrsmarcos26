package main

import "fmt"

func main() {
	var n1 int
	fmt.Println("Me conte um número que eu falarei o antecessor dele e seu sucessor.")
	fmt.Print("Número:")
	fmt.Scan(&n1)
	fmt.Println("Antecessor:", n1-1)
	fmt.Println("Sucessor:", n1+1)
}
