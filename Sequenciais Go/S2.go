package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Me diga um valor de número inteiro e eu te mostro o seu dobro, triplo e quadruplo: ")
	fmt.Scan(&n1)
	// esfregar na cara do usuário
	fmt.Println("Então quer dizer que seu número é:", n1)
	fmt.Println("Sendo assim o seu dobro é:", n1*2)
	fmt.Println("Seu triplo também será:", n1*3)
	fmt.Println("E por fim, o seu quadruplo é:", n1*4)
}
