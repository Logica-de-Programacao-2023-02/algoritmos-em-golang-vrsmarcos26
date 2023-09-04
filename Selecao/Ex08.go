package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("Me diga um número e falarei se ele tem relação com os dias da semana: ")
	fmt.Scan(&n1)
	if n1 == 1 {
		fmt.Print("1 condiz com o dia da semana Domingo")
	} else if n1 == 2 {
		fmt.Print("1 condiz com o dia da semana Segunda")
	} else if n1 == 3 {
		fmt.Print("1 condiz com o dia da semana Terça")
	} else if n1 == 4 {
		fmt.Print("1 condiz com o dia da semana Quarta")
	} else if n1 == 5 {
		fmt.Print("1 condiz com o dia da semana Quinta")
	} else if n1 == 6 {
		fmt.Print("1 condiz com o dia da semana Sexta")
	} else if n1 == 7 {
		fmt.Print("1 condiz com o dia da semana Sábado")
	} else {
		fmt.Print("Seu número não corresponde a um dia da semana.")
	}
}
