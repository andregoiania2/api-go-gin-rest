package main

import (
	"github.com/andregoiania/api-go-gin/database"
	"github.com/andregoiania/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDados()
	routes.HandleRequest()
}
