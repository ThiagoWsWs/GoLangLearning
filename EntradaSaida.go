package main

import "fmt"

func main() {

	var nome string //Declara a varival nome
	teste := "\nTestando Variaveis em GO"

	fmt.Println("Informe seu nome: ") //Preferencia por usar PrintLn pois pula uma linha...
	fmt.Scanf("%s", &nome)            //Armazena o nome
	fmt.Printf("Bem vindo %s!", nome)
	fmt.Printf("%s", teste)
}
