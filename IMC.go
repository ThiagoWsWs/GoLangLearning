package main

import (
	"fmt"
	"math"
)

func main() {
	var peso float64
	var altura float64

	fmt.Printf("Informe seu peso: ")
	fmt.Scanf("%f\n", &peso)
	fmt.Printf("Informe sua altura: ")
	fmt.Scanf("%f\n", &altura)

	imc := peso / (altura * altura)

	fmt.Printf("Seu IMC e de: %.2f", imc)

	//Apresentacao da biblioteca Math

	IMC := peso / float64(math.Pow(altura, 2))
	fmt.Printf("Seu IMC usando a biblioteca Math: %.2f", IMC)
}
