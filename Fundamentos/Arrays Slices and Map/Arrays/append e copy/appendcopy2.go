package main

import "fmt"

func main() {
	barraDeFerramentas := [5]string{"Arquivo", "Usuários", "Documentos", "Finanças", "Ajuda"}
	itensArquivo := []string{"Novo", "Abrir", "Salvar", "Salvar como...", "Exportar", "Fechar"}
	fmt.Println(barraDeFerramentas[0], "\n", itensArquivo[0:6])

	// itensUsuario := make([]string, 4)
	// copy(itensUsuario, itensArquivo)
	// fmt.Println(barraDeFerramentas[1], itensUsuario)
	// itensUsuario = append(itensUsuario, "teste")
	// fmt.Println(itensUsuario)
}
