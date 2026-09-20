package service

import (
	"task-api/internal/model"
	"task-api/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) GetAllUsers() ([]model.User,error) {
	return s.Repo.GetAll()
}

func (s *UserService) GetOneUsers(id int) (*model.User,error) {
	return  s.Repo.GetOne(id)
}

func (s *UserService) PostUser(user_in model.CreateUserRequest) (*model.User,error) {
	return  s.Repo.Post(user_in)
}

func (s *UserService) DeleteUser(id int) (int,error) {
	return s.Repo.Delete(id)
}

func (s *UserService) UpdateUser(user_in model.UpdateUserRequest) (*model.User,error) {
	return s.Repo.Put(user_in)
}
