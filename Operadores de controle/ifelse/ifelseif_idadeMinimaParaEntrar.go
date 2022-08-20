package main

import "fmt"

func idadeMinimaParaEntrar(idade int) string {
	if idade >= 60 {
		return "Vai pra casa véio"
	} else if idade >= 18 && idade < 59 {
		return "Entra doidão"
	} else if idade >= 5 && idade < 18 {
		return "Vaza pirralhada"
	} else {
		return "Digite sua idade corretamente"
	}
}

func main() {
	fmt.Println(idadeMinimaParaEntrar(1))
}
