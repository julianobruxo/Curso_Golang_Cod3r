package main

import "fmt"

func main() {
	outrTeste := [5]int{10, 10, 10, 10, 10}

	mediaDoTeste := 0
	for i := 0; i < len(outrTeste); i++ {
		mediaDoTeste += outrTeste[i]
	}
	result := mediaDoTeste / int(len(outrTeste))
	fmt.Println("Média do teste tem que ser", result)

}
