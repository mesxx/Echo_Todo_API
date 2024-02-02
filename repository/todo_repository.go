package repository

import (
	"github.com/mesxx/Echo_Todo_API/model"

	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (r *TodoRepository) Create(todo *model.Todo) error {
	return r.DB.Create(todo).Error
}

func (r *TodoRepository) FindByID(id uint) (*model.Todo, error) {
	var todo model.Todo
	if err := r.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) Delete(todo *model.Todo) error {
	return r.DB.Delete(todo).Error
}

func (r *TodoRepository) FindAll() ([]model.Todo, error) {
	var todo []model.Todo
	if err := r.DB.Find(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}
