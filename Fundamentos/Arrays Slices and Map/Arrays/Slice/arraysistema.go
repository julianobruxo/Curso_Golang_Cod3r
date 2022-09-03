package main

import "fmt"

func main() {
	barraDeFerramentas := [...]string{"Arquivo", "Editar", "Inserir", "Tipos"}
	fmt.Println(barraDeFerramentas)
	Arquivo := []string{"Novo", "Salvar", "Salvar como...", "Abrir"}
	fmt.Println(barraDeFerramentas[0], Arquivo)
}
