package produtos

import (
	"LearningCrud/database"
	"fmt"
)

func CriarProduto() {
	var nome string
	var preco float64

	fmt.Print("\nDigite o nome do produto: ")
	fmt.Scan(&nome)

	fmt.Print("\nDigite o preço do produto: ")
	fmt.Scan(&preco)

	_, err := database.DB.Exec("INSERT INTO produtos (name, price) VALUES ($1, $2)", nome, preco)
	if err != nil {
		fmt.Println("Erro ao inserir produto:", err)
		return
	}
	fmt.Println("\nProduto cadastrado com sucesso!")
}

func VerProduto() {
	rows, err := database.DB.Query("SELECT id, name, price FROM produtos")
	if err != nil {
		fmt.Println("Erro ao buscar produtos:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n--- Lista de Produtos ---")

	for rows.Next() {
		var id int
		var nome string
		var preco float64

		err := rows.Scan(&id, &nome, &preco)
		if err != nil {
			fmt.Println("Erro ao ler produto:", err)
			continue
		}

		fmt.Printf("ID: %d | Nome: %s | Preço: R$ %.2f\n", id, nome, preco)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Erro ao iterar pelos produtos:", err)
	}
}

func AtualizarProduto() {
	var id int
	fmt.Print("Digite o ID do produto que deseja atualizar: ")
	fmt.Scan(&id)

	var novoNome string
	var novoPreco float64

	fmt.Print("Digite o novo nome do produto: ")
	fmt.Scan(&novoNome)

	fmt.Print("Digite o novo preço do produto: ")
	fmt.Scan(&novoPreco)

	result, err := database.DB.Exec("UPDATE produtos SET name=$1, price=$2 WHERE id=$3", novoNome, novoPreco, id)
	if err != nil {
		fmt.Println("Erro ao atualizar produto:", err)
		return
	}

	linhasAfetadas, _ := result.RowsAffected()
	if linhasAfetadas == 0 {
		fmt.Println("Produto não encontrado")
		return
	}

	fmt.Println("Produto atualizado com sucesso!")
}

func DeletarProduto() {
	var id int
	fmt.Print("Digite o ID do produto que deseja deletar: ")
	fmt.Scan(&id)

	result, err := database.DB.Exec("DELETE FROM produtos WHERE id=$1", id)
	if err != nil {
		fmt.Println("Erro ao deletar produto:", err)
		return
	}

	linhasAfetadas, _ := result.RowsAffected()
	if linhasAfetadas == 0 {
		fmt.Println("Produto não encontrado")
		return
	}

	fmt.Println("Produto deletado com sucesso!")
}
