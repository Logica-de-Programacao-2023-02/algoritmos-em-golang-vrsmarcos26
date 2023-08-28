package main

import "fmt"

func main() {
	var n1 int

	fmt.Print("Me diga um número: ")
	fmt.Scan(&n1)

	if n1%2 == 0 {
		fmt.Print("É par o seguinte número: ", n1)
	} else {
		fmt.Print("É impar o seguinte número: ", n1)
	}

}