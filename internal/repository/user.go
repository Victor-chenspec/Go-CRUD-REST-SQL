package repository

import (
	"context"
	"task-api/internal/model"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	DB *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r * UserRepository) GetAll() ([]model.User,error) {
	rows , err := r.DB.Query(
		context.Background(),
		"SELECT * FROM users ORDER BY id",
	)

	if err != nil {
		return nil , err
	}

	users := []model.User{}

	for rows.Next() {
		var user model.User

		if err := rows.Scan(&user.ID,&user.NAME,&user.EMAIL) ; err != nil {
			return  nil , err
		}

		users = append(users,user)
	}

	return users , nil
}

func (r *UserRepository) GetOne(id int) (*model.User,error) {
	var user model.User

	err := r.DB.QueryRow(
		context.Background(),
		"SELECT * FROM users WHERE id = $1 ",
		id,
	).Scan(&user.ID,&user.NAME,&user.EMAIL)

	if err == pgx.ErrNoRows {
		return nil,nil 
	}

	if err != nil {
		return nil ,err
	}

	return &user , nil
}

func (r *UserRepository) Post(user_in model.CreateUserRequest) (*model.User,error) {
	var user model.User

	err := r.DB.QueryRow(
		context.Background(),
		"INSERT INTO users (name,email) VALUES ($1,$2) RETURNING id , name , email",
		user_in.NAME,
		user_in.EMAIL,
	).Scan(&user.ID,&user.NAME,&user.EMAIL)

	if err != nil {
		return nil ,err
	}

	return &user , nil
}

func (r *UserRepository) Delete(id int) (int,error) {
	result , err := r.DB.Exec(
		context.Background(),
		"DELETE FROM users WHERE id = $1",
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

func (r *UserRepository) Put(user_in model.UpdateUserRequest) (*model.User,error) {
	var user model.User

	err := r.DB.QueryRow(
		context.Background(),
		"UPDATE users SET name = $1 , email = $2 WHERE id = $3 RETURNING id , name , email",
		user_in.NAME,
		user_in.EMAIL,
		user_in.ID,
	).Scan(&user.ID,&user.NAME,&user.EMAIL)

	if err == pgx.ErrNoRows {
		return nil , nil
	}

	if err != nil {
		return nil , err
	}

	return &user , nil
}