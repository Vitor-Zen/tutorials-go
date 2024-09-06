package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cria uma coleção de itens de "diferentes" tipos ou mesmo tipo, nesse caso
type aluno struct {
	ID   string `json:"id"`
	Nome string `json:"nome"`
	RA   string `json:"ra"`
}

// Simula os dados
var alunos = []aluno{
	{ID: "1", Nome: "Lucas bostofre", RA: "203023"},
	{ID: "2", Nome: "Ian arnoide", RA: "202452"},
	{ID: "3", Nome: "Zen deus do bjj", RA: "203100"},
}

func main() {
	// Criação do servidor
	router := gin.Default()

	// Criação das rotas
	router.GET("/alunos", getAlunos)
	router.GET("/alunos/:id", getAlunoByID)
	router.POST("/alunos", postAlunos)

	router.Run("localhost:8080")
}

// c*gin.Context nada mais é do que um ponteiro para Context(parte mais importante do GIN)
func getAlunos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, alunos) // reponsavel por serializar nossa resposta em formato JSON
}

func postAlunos(c *gin.Context) {
	var newAluno aluno // Cria variavel newAluno que é do tipo aluno (struct criado no topo do código)

	if err := c.BindJSON(&newAluno); err != nil { // 	// Tenta pegar os dados da requisição e coloca esses dados em newAluno; Caso não de certo, da um return (não faz nada)
		return
	}

	alunos = append(alunos, newAluno)            // Adiciona o newAluno no final da lista alunos
	c.IndentedJSON(http.StatusCreated, newAluno) // Envia de volta uma resposta em formato JSON
}

func getAlunoByID(c *gin.Context) {
	id := c.Param("id") // Pega o ID da requisição

	for _, a := range alunos { // Ignora o indice e checa cada elemento do slice alunos
		if a.ID == id { // Se o id passado for igual ao id do elemento,
			c.IndentedJSON(http.StatusOK, a) // retorna o objeto a no formato JSON
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "alunos não encontrados"})
}
