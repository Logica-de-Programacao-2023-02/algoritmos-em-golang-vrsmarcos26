package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Qual sua idade? R:")
	fmt.Scan(&n1)
	fmt.Print("Essa é sua idade em dias: ", n1*365)
}
