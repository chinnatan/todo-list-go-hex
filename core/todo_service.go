package core

import "errors"

type TodoService interface {
	CreateTodo(todo Todo) error
	GetAll() ([]Todo, error)
	GetById(id int64) (Todo, error)
}

type todoServiceImpl struct {
	repo TodoRepository
}

func NewTodoService(repo TodoRepository) TodoService {
	return &todoServiceImpl{repo: repo}
}

func (s *todoServiceImpl) CreateTodo(todo Todo) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}
	if err := s.repo.Save(todo); err != nil {
		return err
	}
	return nil
}

func (s *todoServiceImpl) GetAll() ([]Todo, error) {
	items, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *todoServiceImpl) GetById(id int64) (Todo, error) {
	item, err := s.repo.FindById(id)
	if err != nil {
		return Todo{}, err
	}
	return item, nil
}
