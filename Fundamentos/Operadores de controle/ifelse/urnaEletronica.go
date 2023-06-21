package main

import (
	"fmt"
)

func msgUrnaEltronica(nCandidato int) string {
	if nCandidato == 17 {
		return "Lá vamos nós de novo..."
	} else if nCandidato == 13 {
		return "Digite novamente até acertar"
	} else if nCandidato == 24 {
		return "Gosta de defumar linguiça a peido, né?"
	} else {
		return "Escolha uma das 3 opções"
	}
}

func main() {
	fmt.Println(msgUrnaEltronica(24))
}
