package service

import (
	todo "github.com/AndreyDuda/go-api-todo"
	"github.com/AndreyDuda/go-api-todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Todolist interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId int, listId int) (todo.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	Todolist
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAutService(repos.Authorization),
		Todolist:      NewTodoListService(repos.Todolist),
	}
}
