package main

import "fmt"

func main() {

	str := "\n "
	fmt.Printf("%s", str)

	str = "\nAlterado\n"
	fmt.Printf("%s", str)

	frase := "Indice"                                    //6 indices (caracteres)
	PrimeiroCarc := frase[0]                             // puxa o primeiro caractere, no indice 0
	IndiceUltimoCarc := len(frase) - 1                   //Len conta todo os caracteres da string, -1 para achar a ultima letra...
	UltimoCarc := frase[IndiceUltimoCarc]                //Recebe o caractere da strig Frase na posicao IndiceUltimoCarc
	fmt.Printf("Primeiro Caractere: %c\n", PrimeiroCarc) //Informa o primeiro caractere (indice na posicao 0)
	fmt.Printf("Ultimo Caractere: %c \n", UltimoCarc)    //Informa caractere na posicao IndiceUltimoCarc

	nome := "Thiago Winicius da Silva"                      //Uma string definida
	tamanhoFrase := len(nome)                               //Conta caracteres dessa string
	fmt.Printf("o nome tem %d caracteres \n", tamanhoFrase) //Pega o tamanho da frase e apresenta
	primeiraPalavra := nome[0:5]                            //Recebe os indices das palavras que vai apresentar
	segundaPalavra := nome[7:15]
	terceiraPalavra := nome[16:24]
	fmt.Printf("A primeira Palavra eh: %s \n", primeiraPalavra) //Informa a primeira palavra baseado nos indices recibos anteriormente
	fmt.Printf("A segunda palavra eh: %s\n", segundaPalavra)    //Informa a segunda palavra
	fmt.Printf("A terceira palavra eh: %s\n", terceiraPalavra)  //Informa a terceira palavra
}
