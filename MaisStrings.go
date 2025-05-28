package main

import (
	"fmt"
	"strings"
)

func main() {
	frase := "Bolo de Chocolate"
	var palavra string
	fmt.Printf("Frase Original: %s\n", frase)
	fraseMinuscula := strings.ToLower(frase) // Transforma toda a frase para minusculo
	fmt.Printf("Frase Minuscula: %s\n", fraseMinuscula)
	fraseMaiuscula := strings.ToUpper(frase) // Transforma toda a frase para maiusculo
	fmt.Printf("Frase Maiuscula: %s\n", fraseMaiuscula)

	frase2 := "   Baita Frase    " //Frase cheia de espacos
	fmt.Printf("A frase2 original: %s\n", frase2)
	tamanhofrase2 := len(frase2) //Conta a frase (todos os espacos inclusive)
	fmt.Printf("O tamanho da frase2 eh: %d\n", tamanhofrase2)
	frase2 = strings.Trim(frase2, " ")                               //Remove todos os espacos da frase
	tamanhofrase2 = len(frase2)                                      //Conta a frase novamente
	fmt.Printf("O tamanho da frase 2 agora eh: %d\n", tamanhofrase2) //Apresenta a nova frase sem espacos
	fmt.Printf("A frase agora eh: %s\n", frase2)                     //Apresenta a nova frase sem espacos

	fmt.Println("Digite um nova palavra: ")
	fmt.Scanf("%s", &palavra)
	tamanhoPalavra := len(palavra)                                        //Conta caracteres dela
	metade := tamanhoPalavra / 2                                          //Divide esses caracteres por 2
	fmt.Printf("Primeira metade da palavra eh: %s \n", palavra[0:metade]) //Apresenta a primeira metade da palavra
	fmt.Printf("Segunda metada da palavra eh: %s \n", palavra[metade:])   //Apresenta a segunda metade da palavra

}
