package service 

import (

	"latihan_4/model"
	"latihan_4/repository"

)


type TodoService struct {

	repo *repository.TodoRepository

}

func NewTodoService(r *repository.TodoRepository) *TodoService {
	return &TodoService{repo:r}
}

func (s *TodoService) AddTodo(title string) model.Todo{

	todo := model.Todo{
		Title: title, 
		Done: false,
	}

	return s.repo.Save(todo)
}

func (s *TodoService) GetTodos() []model.Todo {
	return s.repo.FindAll()
}