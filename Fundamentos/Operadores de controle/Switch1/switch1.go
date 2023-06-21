package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "Congrats modafaca"
	case 8, 7:
		return "Almost there dude"
	case 6, 5:
		return "Pick up the pace"
	case 4, 3:
		return "WTF r u doing?"
	case 2, 1, 0:
		return "Beat it dumbass"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(6))
}
