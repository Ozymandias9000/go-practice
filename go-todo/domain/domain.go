package domain

type UserRepo interface {
	GetByID(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	CreateUser(user *User) (*User, error)
}

type TodoRepo interface {
	GetAllTodos() ([]*Todo, error)
	GetByID(id int64) (*Todo, error)
	CreateTodo(todo *Todo) (*Todo, error)
	UpdateTodo(todo *Todo) (*Todo, error)
	DeleteTodo(todo *Todo) (*Todo, error)
}

type DB struct {
	UserRepo UserRepo
	TodoRepo TodoRepo
}

type Domain struct {
	DB DB
}
