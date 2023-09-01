package routes

import (
	"github.com/andregoiania/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("alunos/cpf/:cpf", controllers.BuscaAlunoPorCPf)
	r.Run()
}
