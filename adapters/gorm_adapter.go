package adapters

import (
	"todo-list/core"

	"gorm.io/gorm"
)

type gormTodoRepository struct {
	db *gorm.DB
}

func NewGormTodoRepository(db *gorm.DB) *gormTodoRepository {
	return &gormTodoRepository{db: db}
}

func (r *gormTodoRepository) Save(todo core.Todo) error {
	if result := r.db.Create(&todo); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *gormTodoRepository) FindAll() ([]core.Todo, error) {
	var todos []core.Todo
	if result := r.db.Find(&todos); result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

func (r *gormTodoRepository) FindById(id int64) (core.Todo, error) {
	var todo core.Todo
	if result := r.db.First(&todo, id); result.Error != nil {
		return core.Todo{}, result.Error
	}
	return todo, nil
}
