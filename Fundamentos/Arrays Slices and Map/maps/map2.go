package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Juliano":  15300.25,
		"Suzy":     20500.50,
		"Isabelle": 10520.50,
	}
	funcsESalarios["Joao Castanha"] = 50500.25
	fmt.Println(funcsESalarios)
	funcsESalarios["Juliano"] = 12300.50

	for nome, valor := range funcsESalarios {
		fmt.Println(nome, valor)
	}
}
