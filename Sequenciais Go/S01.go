package main

import "fmt"

func main() {
	//Abrindo variável
	var n1, n2, n3 float64
	//Conversa com o usuário
	fmt.Println("Vamos fazer uma soma de três números.")
	fmt.Print("Me fale o primeiro número que você queira somar: ")
	fmt.Scan(&n1)
	fmt.Print("Me fale o segundo número que você queira somar: ")
	fmt.Scan(&n2)
	fmt.Print("Me fale o terceiro número que você queira somar: ")
	fmt.Scan(&n3)
	//Resultado
	fmt.Print("O resultado da soma dos números é igual a: ", n1+n2+n3)
}
