package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"latihan_4/service"

)

type Handler struct{
	service *service.TodoService
	reader *bufio.Reader
}

func NewHandler(s *service.TodoService) *Handler {
	return &Handler{
		service: s,
		reader: bufio.NewReader(os.Stdin),
	}
                     
}

func (h *Handler) Run(){

	for {

		fmt.Println("\n1. Add Todo")
		fmt.Println("\n2. List Todo ")
		fmt.Println("\n3. Exit ")
		fmt.Print("Choose: ")

		input, _ := h.reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice , err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid Input")
			continue
		}
		
		switch choice {
			case 1 : 
				h.addTodo()
			case 2 : 
				h.listTodo()
			case 3 :
				fmt.Println("Bye")
				return
			default : 
				fmt.Println("Invalid Choice")
		}
	}
}

func (h *Handler) addTodo() {

	fmt.Print("Enter title: ")
	
	title, _ := h.reader.ReadString('\n')
	title = strings.TrimSpace(title)

	if title == "" {

		fmt.Println("Title cannot be empty")
		return
	}

	todo := h.service.AddTodo(title)
	fmt.Println("Added:", todo)
}

func (h * Handler) listTodo() {

	todos := h.service.GetTodos()
	for _, t := range todos {
		fmt.Println(t.ID,"-",t.Title,"-",t.Done)
	}

}