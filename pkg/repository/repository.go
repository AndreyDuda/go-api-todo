package repository

type Authorization interface {
}

type Todolist interface {
}

type TodoItem interface {
}

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}
