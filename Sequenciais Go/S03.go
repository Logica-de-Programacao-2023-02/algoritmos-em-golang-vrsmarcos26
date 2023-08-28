package main

import "fmt"

func main() {
	var p1, a1 float64
	fmt.Println("Para você poder monitorar sua saúde melhor, vamos calcular seu IMC.")
	fmt.Println("Para isso precisarei que você me diga")
	fmt.Print("Peso:")
	fmt.Scan(&p1)
	fmt.Print("Altura:")
	fmt.Scan(&a1)
	//Calculo aqui pq pc é burro e não faz tudo em uma linha simples
	S := a1 * a1
	fmt.Print("Com esses dados, podemos dizer que seu IMC é: ", p1/S)
}
