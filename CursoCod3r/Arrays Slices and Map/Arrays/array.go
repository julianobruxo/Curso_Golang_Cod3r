package main

import "fmt"

func main() {
	// arrrays são estruturas estáticas e fixas, ou seja, com elementos do mesmo tipo sempoe e sempre do mesmo tamanho
	var notas [3]int
	fmt.Println(notas)
	notas[0], notas[1], notas[2] = 4, 5, 6
	fmt.Println(notas)

	totalDasNotasDoArray := 0
	for l := 0; l < len(notas); l++ {
		totalDasNotasDoArray += notas[l]
	}
	fmt.Println(totalDasNotasDoArray)
	media := totalDasNotasDoArray / int(len(notas))
	fmt.Print("Média final ", media)

}
