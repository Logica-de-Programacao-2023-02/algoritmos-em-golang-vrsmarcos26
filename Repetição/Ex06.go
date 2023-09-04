package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Me fale um número para eu le mestrar a tabuada do mesmo R: ")
	fmt.Scan(&n1)
	for i := 1; i <= 10; i++ {
		r := n1 * i
		fmt.Printf("%d x %d = %d\n", n1, i, r)
	}
}
