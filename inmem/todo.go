package inmem

import (
	"context"
	"sync"
	"todo"
)

type Service struct {
	nextID int
	mu     sync.Mutex
	todos  []*todo.Todo
}

func NewService() todo.Service {
	return &Service{
		nextID: 1,
		todos: make([]*todo.Todo, 0),
	}
}

func (s *Service) CreateTodo(ctx context.Context, request todo.CreateTodoRequest) (*todo.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t := &todo.Todo{
		ID:       s.nextID,
		Value:    request.Value,
		Complete: false,
	}
	s.todos = append(s.todos, t)
	s.nextID++

	return t, nil
}

func (s *Service) UpdateTodo(ctx context.Context, request todo.UpdateTodoRequest) (*todo.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, err := s.getTodoByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}
	t.Value = request.Value
	t.Complete = request.Complete

	return t, nil
}

func (s *Service) DeleteTodo(ctx context.Context, request todo.DeleteTodoRequest) error {
	panic("implement me")
}

func (s *Service) GetTodoByID(ctx context.Context, request todo.GetTodoByIDRequest) (*todo.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, err := s.getTodoByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *Service) GetAllTodos(ctx context.Context) ([]*todo.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.todos, nil
}

func (s *Service) getTodoByID(_ context.Context, id int) (*todo.Todo, error) {
	for i := range s.todos {
		if s.todos[i].ID == id {
			return s.todos[i], nil
		}
	}
	return nil, todo.Errorf(todo.ENOTFOUND, "Todo with ID '%d' could not be found.", id)
}
