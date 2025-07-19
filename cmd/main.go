package main

import (
	"LearningCrud/database"
	"LearningCrud/produtos"
	"fmt"
)

func main() {
	database.Conectar()
	defer database.DB.Close()

	for {
		var opcao int
		fmt.Println("\nMenu de Produtos")
		fmt.Print("\n1. Criar produto")
		fmt.Print("\n2. Ver lista de produtos")
		fmt.Print("\n3. Atualizar produto")
		fmt.Print("\n4. Deletar um produto")
		fmt.Print("\n0. Sair")
		fmt.Print("\nEscolha uma opcão: ")
		fmt.Scan(&opcao)

		switch opcao {
		case 0:
			fmt.Println("Volte sempre!")
			return
		case 1:
			produtos.CriarProduto()
		case 2:
			produtos.VerProduto()
		case 3:
			produtos.AtualizarProduto()
		case 4:
			produtos.DeletarProduto()
		}
	}
}
