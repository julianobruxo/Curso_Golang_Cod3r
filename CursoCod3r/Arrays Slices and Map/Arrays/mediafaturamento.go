package main

import "fmt"

func main() {
	faturamentoMensal := [6]int{10, 12, 15, 12, 13, 18}

	MediaSemestre := 0
	for a := 0; a < len(faturamentoMensal); a++ {
		MediaSemestre += faturamentoMensal[a]
	}
	result := MediaSemestre / int(len(faturamentoMensal))
	fmt.Println("Faturamento médio do Primeiro semestre", result)

}
