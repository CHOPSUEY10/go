package repository

import "latihan_4/model"

type TodoRepository struct {

	todos []model.Todo
	nextID int

}

func NewTodoRepository() *TodoRepository  {

	return &TodoRepository{
		todos: []model.Todo{},
		nextID: 1,
	}
}

func (r *TodoRepository) Save(todo model.Todo) model.Todo {

	todo.ID = r.nextID
	r.nextID++
	r.todos = append(r.todos,todo)
	return todo

}

func (r *TodoRepository) FindAll() []model.Todo{
	return r.todos
}