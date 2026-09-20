package service

import (
	"task-api/internal/model"
	"task-api/internal/repository"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{
		Repo: repo,
	}
}

func (s *TaskService) GetAll(id int) ([]model.Task,error) {
	return s.Repo.GetAll(id)
}

func (s *TaskService) PostTask(task_in model.CreateTaskRequest) (*model.Task,error) {
	return s.Repo.Post(task_in)
}

func (s *TaskService) UpdateTask(task_in model.UpdateTaskRequest) (*model.Task,error) {
	return s.Repo.Put(task_in)
}

func (s *TaskService) DeleteTask(id int) (int , error) {
	return s.Repo.Delete(id)
}