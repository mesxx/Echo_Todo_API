package usecase

import (
	"github.com/mesxx/Echo_Todo_API/model"
	"github.com/mesxx/Echo_Todo_API/repository"
)

type TodoUsecase struct {
	TodoRepository repository.TodoRepository
}

func NewTodoUsecase(tu repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{TodoRepository: tu}
}

func (u *TodoUsecase) Create(todo *model.Todo) error {
	return u.TodoRepository.Create(todo)
}

func (u *TodoUsecase) Delete(todo *model.Todo) error {
	return u.TodoRepository.Delete(todo)
}

func (u *TodoUsecase) FindAll() ([]model.Todo, error) {
	return u.TodoRepository.FindAll()
}

func (u *TodoUsecase) FindByID(id uint) (*model.Todo, error) {
	return u.TodoRepository.FindByID(id)
}
