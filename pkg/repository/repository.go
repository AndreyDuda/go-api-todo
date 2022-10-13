package repository

import (
	todo "github.com/AndreyDuda/go-api-todo"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username string, password string) (todo.User, error)
}

type Todolist interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId int, listId int) (todo.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId int, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId int, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId int, itemId int, input todo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	Todolist
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Todolist:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
