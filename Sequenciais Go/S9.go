package main

import "fmt"

func main() {
	var n1 float64
	fmt.Print("Qual valor do item que você deseja desconto de 10%? R:")
	fmt.Scan(&n1)
	fmt.Print("Com o desconto, você pagara nesse item: ", n1*0.90)
}
