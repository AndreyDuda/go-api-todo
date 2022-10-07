package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Todolist interface {
}

type TodoItem interface {
}

type Repository struct {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
