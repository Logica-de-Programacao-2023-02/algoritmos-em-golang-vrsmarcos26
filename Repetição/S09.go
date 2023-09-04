package main

import (
	"fmt"
)

func main() {
	var s, q, n1 int
	n1 = -1
	for n1 != 0 {

		fmt.Println("Digite um número: ")
		fmt.Scan(&n1)
		s += n1
		q++
	}
	media := s / q
	fmt.Print("A média é ", media)

}
