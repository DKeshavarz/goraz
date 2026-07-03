package repository

import (
	"database/sql"

	"github.com/DKeshavarz/goraz/module3/taskmanger/model"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) FetchAll() ([]model.Task, error) {
	// TODO:
	return nil, nil
}
func (r *PostgresRepo) Create(task model.Task) error {
	// TODO:
	return nil
}

func (r *PostgresRepo) GetByID(id int) (*model.Task, error) {
	// TODO:
	return nil, nil
}

func (r *PostgresRepo) Delete(id int) error {
	// TODO:
	return nil
}
