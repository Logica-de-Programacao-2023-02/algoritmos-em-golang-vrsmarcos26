package main

import "fmt"

func main() {
	var n1, n2, n3 float64
	fmt.Print("Me fale um número: ")
	fmt.Scan(&n1)
	fmt.Print("Me fale um número: ")
	fmt.Scan(&n2)
	fmt.Print("Me fale um número: ")
	fmt.Scan(&n3)

	if n1 <= n2 && n1 <= n3 {
		fmt.Printf("%.2f, ", n1)
		if n2 <= n3 {
			fmt.Printf("%.2f, %.2f\n", n2, n3)
		} else {
			fmt.Printf("%.2f, %.2f\n", n3, n2)
		}
	} else if n2 <= n1 && n2 <= n3 {
		fmt.Printf("%.2f, ", n2)
		if n1 <= n3 {
			fmt.Printf("%.2f, %.2f\n", n1, n3)
		} else {
			fmt.Printf("%.2f, %.2f\n", n3, n1)
		}
	} else {
		fmt.Printf("%.2f, ", n3)
		if n1 <= n2 {
			fmt.Printf("%.2f, %.2f\n", n1, n2)
		} else {
			fmt.Printf("%.2f, %.2f\n", n2, n1)
		}
	}
}
