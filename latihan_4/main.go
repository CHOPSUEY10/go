package main

import (

	"latihan_4/cli"
	"latihan_4/repository"
	"latihan_4/service"

)

func main () {

	repo := repository.NewTodoRepository()
	svc := service.NewTodoService(repo)
	handler := cli.NewHandler(svc)

	handler.Run() 
}