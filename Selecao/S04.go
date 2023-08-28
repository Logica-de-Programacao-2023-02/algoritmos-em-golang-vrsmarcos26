package main

import "fmt"

func main() {
	var altura float32
	var peso float32

	fmt.Print("Me indique sua altura: ")
	fmt.Scan(&altura)
	fmt.Print("Me indique seu peso: ")
	fmt.Scan(&peso)

	imc := peso / (altura * altura)

	if imc >= 18.5 && imc <= 24.9 {
		fmt.Print("O seu imc está dentro do ideal.")
	}
	if imc > 18.5 && imc < 24.9 {
		fmt.Print("O seu imc está acima do ideal.")
	}
	if imc < 18.5 && imc < 24.9 {
		fmt.Print("O seu imc está abaixo do ideal.")
	}
}
