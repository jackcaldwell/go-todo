package todo

import "context"

type Service interface {
	CreateTodo(ctx context.Context, request CreateTodoRequest) (*Todo, error)
	UpdateTodo(ctx context.Context, request UpdateTodoRequest) (*Todo, error)
	DeleteTodo(ctx context.Context, request DeleteTodoRequest) error
	GetTodoByID(ctx context.Context, request GetTodoByIDRequest) (*Todo, error)
	GetAllTodos(ctx context.Context) ([]*Todo, error)
}

// Middleware describes a service (as opposed to endpoint) middleware for the Service.
type Middleware func(service Service) Service

type CreateTodoRequest struct {
	Value    string `json:"value"`
	Complete bool   `json:"complete"`
}

type UpdateTodoRequest struct {
	ID       int    `json:"id"`
	Value    string `json:"value"`
	Complete bool   `json:"complete"`
}

type DeleteTodoRequest struct {
	ID int `json:"id"`
}

type GetTodoByIDRequest struct {
	ID int `json:"id"`
}

type Todo struct {
	ID       int    `json:"id"`
	Value    string `json:"value"`
	Complete bool   `json:"complete"`
}
