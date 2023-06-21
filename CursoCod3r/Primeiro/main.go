package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI = 3.1415
	var raio = 3.2 // forma completa de se criar variáveis em GO

	area := PI * m.Pow(raio, 2) // Posso usar m do import ao invés de Math
	fmt.Println("A área da circunferência é ", area)

	const (
		a      = ((10 * 5 * 2) / 2)
		result = a
	)

	fmt.Print(result)
}

/* Oque eu aprendi nessa aula?
Aprendi que consts e vars sempre devem ser usadas, não podendo apenas serem declaradas;
Aprendi que posso referenciar um import a partir de uma unica letra
Que é possível colocar várias const ou vars dentro de uma const ou var aberta por ()
Que eu não preciso usar o tipo do numero ou valor atribuído ie. "float64" pois o compilador percebe isso*/
