package main

import "fmt"

func main() {
	var n1, n2, n3 float64
	fmt.Println("Vamos calcular média ponderada de três números a sua escolha om pesos 2, 3 e 5.")
	fmt.Print("Primeiro número:")
	fmt.Scan(&n1)
	fmt.Print("Segundo número:")
	fmt.Scan(&n2)
	fmt.Print("Terceiro número:")
	fmt.Scan(&n3)
	//INFELIZMENTE tem que fzr calculo aqui pq pc é burro
	s1 := n1 * 2
	s2 := n2 * 3
	s3 := n3 * 5
	soma := s1 + s2 + s3
	fmt.Print("A média ponderada desses três números é: ", soma/10)
}
