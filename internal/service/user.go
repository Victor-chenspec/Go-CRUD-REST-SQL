package service

import (
	"task-api/apperror"
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
	users , err := s.Repo.GetAll()

	if err != nil {
		return nil , apperror.InternalError("Database error")
	}

	return users , nil
}

func (s *UserService) GetOneUsers(id int) (*model.User,error) {
	user ,err := s.Repo.GetOne(id)

	if err != nil {
		return nil , apperror.InternalError("Database error")
	}

	if user == nil {
		return nil , apperror.NotFound("User not found")
	}

	return  user , nil
}

func (s *UserService) PostUser(user_in model.CreateUserRequest) (*model.User,error) {
	user , err :=  s.Repo.Post(user_in)

	if err != nil {
		return  nil , apperror.InternalError("Database error")
	}

	return  user , nil
}

func (s *UserService) DeleteUser(id int) (int,error) {
	row_affected , err := s.Repo.Delete(id)

	if err != nil {
		return  0 , apperror.InternalError("Database error")
	}

	if row_affected == 0 {
		return  0 , apperror.NotFound("User not found")
	}

	return 1 , nil 
}

func (s *UserService) UpdateUser(user_in model.UpdateUserRequest) (*model.User,error) {
	user , err := s.Repo.Put(user_in)

	if err != nil {
		return nil , apperror.InternalError("Database error")
	}

	if user == nil {
		return  nil , apperror.NotFound("User not found")
	}
	
	return  user ,err
}
