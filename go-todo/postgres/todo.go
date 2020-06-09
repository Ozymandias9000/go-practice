package postgres

import (
	"errors"
	"go-todo/domain"

	"github.com/go-pg/pg/v9"
)

type TodoRepo struct {
	DB *pg.DB
}

func NewTodoRepo(DB *pg.DB) *TodoRepo {
	return &TodoRepo{DB}
}

func (t TodoRepo) GetByID(id int64) (*domain.Todo, error) {
	todo := new(domain.Todo)

	err := t.DB.Model(todo).Where("id = ?", id).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}

	return todo, nil
}

func (t TodoRepo) CreateTodo(todo *domain.Todo) (*domain.Todo, error) {
	_, err := t.DB.Model(todo).Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t TodoRepo) UpdateTodo(todo *domain.Todo) (*domain.Todo, error) {
	_, err := t.DB.Model(todo).Where(" id = ? ", todo.ID).Returning("*").Update()
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t TodoRepo) DeleteTodo(todo *domain.Todo) (*domain.Todo, error) {
	_, err := t.DB.Model(todo).Where(" id = ? ", todo.ID).Delete()
	if err != nil {
		return nil, err
	}

	return todo, nil
}
