package repository

import (
	"context"
	"task-api/internal/model"

	"github.com/jackc/pgx/v5"
)

type TaskRepository struct {
	DB *pgx.Conn
}

func NewTaskRepository(db *pgx.Conn) *TaskRepository {
	return &TaskRepository{
		DB: db,
	}
}

func (r *TaskRepository) GetAll(id int) ([]model.Task,error) {
	rows , err := r.DB.Query(
		context.Background(),
		"SELECT * FROM tasks WHERE user_id = $1",
		id,
	)

	if err != nil {
		return nil , err
	}

	tasks := []model.Task{}

	for rows.Next() {
		var task model.Task

		if row_err := rows.Scan(&task.ID,&task.TITLE,&task.COMPLETED,&task.USERID) ; row_err != nil {
			return nil , row_err
		}

		tasks = append(tasks, task)
	}

	return tasks , nil
}

func (r *TaskRepository) Post(task_in model.CreateTaskRequest) (*model.Task,error){
	var task model.Task

	err := r.DB.QueryRow(
		context.Background(),
		"INSERT INTO tasks (title,user_id) VALUES ($1,$2) RETURNING id , title , completed , user_id",
		task_in.TITLE,
		task_in.USERID,
	).Scan(&task.ID,&task.TITLE,&task.COMPLETED,&task.USERID)

	if err != nil {
		return nil , err
	}

	return  &task , nil
}

func (r *TaskRepository) Put(task_in model.UpdateTaskRequest) (*model.Task,error) {
	var task model.Task

	err := r.DB.QueryRow(
		context.Background(),
		"UPDATE tasks SET title = $1 , completed = $2 , user_id = $3 WHERE id = $4 RETURNING id,title,completed,user_id",
		task_in.TITLE,
		task_in.COMPLETED,
		task_in.USERID,
		task_in.ID,
	).Scan(&task.ID,&task.TITLE,&task.COMPLETED,&task.USERID)

	if err == pgx.ErrNoRows {
		return nil , nil
	}

	if err != nil {
		return nil , err
	}

	return &task ,nil
}

func (r *TaskRepository) Delete(id int) (int , error) {

	result , err := r.DB.Exec(
		context.Background(),
		"DELETE FROM tasks WHERE id = $1",
		id,
	)

	if err != nil {
		return 0 , err
	}

	if result.RowsAffected() == 0 {
		return 0 , nil
	}

	return 1 , nil
}

