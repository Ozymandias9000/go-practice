package domain

import (
	"time"
)

type Todo struct {
	ID        int64  `json:"id"`
	Body      string `json:"body"`
	UserID    int64  `json:"user_id"`
	Completed bool   `json:"completed"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateTodoPayload struct {
	Body string `json:"body" validate:"required"`
}

func (d *Domain) CreateTodo(payload *CreateTodoPayload, currentUserID int64) (*Todo, error) {
	data := &Todo{
		Body:   payload.Body,
		UserID: currentUserID,
	}

	todo, err := d.DB.TodoRepo.CreateTodo(data)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (d *Domain) GetTodoByID(id int64) (*Todo, error) {
	todo, err := d.DB.TodoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

type UpdateTodoPayload struct {
	Body      string `json:"body" `
	Completed *bool  `json:"completed"`
}

func (d *Domain) UpdateTodo(payload UpdateTodoPayload, todo Todo) (*Todo, error) {
	updated := false

	if payload.Body != "" {
		todo.Body = payload.Body
		updated = true
	}

	if payload.Completed != nil {
		todo.Completed = *payload.Completed
		updated = true
	}

	if updated {
		todo.UpdatedAt = time.Now()
	}

	updatedTodo, err := d.DB.TodoRepo.UpdateTodo(&todo)
	if err != nil {
		return nil, err
	}

	return updatedTodo, nil
}

func (d *Domain) DeleteTodo(todo Todo) (*Todo, error) {
	deletedTodo, err := d.DB.TodoRepo.DeleteTodo(&todo)
	if err != nil {
		return nil, err
	}

	return deletedTodo, nil
}
