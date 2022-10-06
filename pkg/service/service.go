package service

import "github.com/AndreyDuda/go-api-todo/pkg/repository"

type Authorization interface {
}

type Todolist interface {
}

type TodoItem interface {
}

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
