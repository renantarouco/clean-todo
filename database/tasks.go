package database

import (
	"database/sql"
	"fmt"

	// Import of the postgres driver to use with driver/sql.
	_ "github.com/lib/pq"

	"github.com/renantarouco/clean-todo/errors"
	"github.com/renantarouco/clean-todo/model"
)

type TasksDB struct {
	conn *sql.DB
}

func NewTasksDB(addr string) (*TasksDB, error) {
	connString := fmt.Sprintf("postgresql://%s?sslmode=disable", addr)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return &TasksDB{
		conn: db,
	}, nil
}

func (db *TasksDB) Put(model.Task) error {
	return errors.ErrNotImplementYet
}

func (db *TasksDB) All() ([]model.Task, error) {
	return nil, errors.ErrNotImplementYet
}
