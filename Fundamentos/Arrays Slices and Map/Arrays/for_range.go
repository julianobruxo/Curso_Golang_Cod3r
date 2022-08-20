package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 10, 5} // elipsis serve para dizer ao compilador contar o numero exato de itens dentro do array
	for i, numero := range arr {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, num := range arr {
		fmt.Println(num)
	}
}
