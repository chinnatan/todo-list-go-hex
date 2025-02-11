package core

type TodoRepository interface {
	Save(todo Todo) error
	FindAll() ([]Todo, error)
	FindById(id int64) (Todo, error)
}
