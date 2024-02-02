package usecase

import (
	"github.com/mesxx/Echo_Todo_API/model"
	"github.com/mesxx/Echo_Todo_API/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepository: ur}
}

func (u *UserUsecase) Create(user *model.User) error {
	return u.UserRepository.Create(user)
}

func (u *UserUsecase) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (u *UserUsecase) CheckPasswordHash(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *UserUsecase) FindAll() ([]model.User, error) {
	return u.UserRepository.FindAll()
}

func (u *UserUsecase) FindByID(id uint) (*model.User, error) {
	return u.UserRepository.FindByID(id)
}

func (u *UserUsecase) Update(user *model.User) error {
	return u.UserRepository.Update(user)
}

func (u *UserUsecase) Delete(id uint) error {
	user, err := u.UserRepository.FindByID(id)
	if err != nil {
		return err
	}
	return u.UserRepository.Delete(user)
}

func (u *UserUsecase) FindByUsername(username string) (*model.User, error) {
	return u.UserRepository.FindByUsername(username)
}
