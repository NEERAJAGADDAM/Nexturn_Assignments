package repository

import (
	"blog-management/model"
	"database/sql"
	"errors"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(user *model.User) error {
	_, err := repo.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	row := repo.DB.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}
