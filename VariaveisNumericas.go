package main

import (
	"fmt"
)

func main() {
	quantidade := 34
	temperatura := -42.23
	peso := 68.5
	var altura float32
	var palavra string
	var tamanho int

	fmt.Printf("A quantidade eh: %d e a temperatura eh: %f\n", quantidade, temperatura)
	fmt.Printf("Thiago pesa: %.2f\n", peso)

	fmt.Printf("Qual sua altura Thiago? \n")
	fmt.Scanf("%f", &altura)
	fmt.Printf("Que legal, voce tem %.2f de altura\n", altura)

	fmt.Printf("Digite uma palavra aleatoria: \n")
	fmt.Scanf("%s", &palavra)
	tamanho = len(palavra)
	fmt.Printf("Essa palavra tem %d caracteres\n", tamanho)
}
