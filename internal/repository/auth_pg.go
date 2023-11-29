package repository

import (
	"fmt"

	"github.com/43nvy/todo-app"
	"github.com/jmoiron/sqlx"
)

type AuthPG struct {
	db *sqlx.DB
}

func NewAuthPG(db *sqlx.DB) *AuthPG {
	return &AuthPG{db: db}
}

func (a *AuthPG) CreateUser(user todo.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := a.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
