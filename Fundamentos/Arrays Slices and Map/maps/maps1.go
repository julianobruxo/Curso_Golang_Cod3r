package main

import "fmt"

func main() {
	aprovados := make(map[int]string)
	aprovados[12345] = "Juliano"
	aprovados[6789] = "Suzy"
	aprovados[101112] = "Joao"
	//fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF:%d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345])
	delete(aprovados, 111213)
	fmt.Println(aprovados[111213])
}
