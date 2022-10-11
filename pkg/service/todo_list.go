package service

import (
	todo "github.com/AndreyDuda/go-api-todo"
	"github.com/AndreyDuda/go-api-todo/pkg/repository"
)

type TodoListService struct {
	repo repository.Todolist
}

func NewTodoListService(repo repository.Todolist) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId int, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}
