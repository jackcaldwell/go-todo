package http

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"todo"
)

func (s *Server) configureHandlers() {
	e := MakeServerEndpoints(s.TodoService)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(s.Logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	s.router.Handle(
		"/api/todos",
		httptransport.NewServer(
			e.CreateTodoEndpoint,
			decodeCreateTodoRequest,
			encodeResponse,
			options...,
		),
	).Methods("POST")

	s.router.Handle(
		"/api/todos/{id}",
		httptransport.NewServer(
			e.UpdateTodoEndpoint,
			decodeUpdateTodoRequest,
			encodeResponse,
			options...,
		),
	).Methods("PUT")

	s.router.Handle(
		"/api/todos/{id}",
		httptransport.NewServer(
			e.DeleteTodoEndpoint,
			decodeDeleteTodoRequest,
			encodeResponse,
			options...,
		),
	).Methods("DELETE")

	s.router.Handle(
		"/api/todos/{id}",
		httptransport.NewServer(
			e.GetTodoByIDEndpoint,
			decodeGetTodoByIDRequest,
			encodeResponse,
			options...,
		),
	).Methods("GET")

	s.router.Handle(
		"/api/todos",
		httptransport.NewServer(
			e.GetAllTodosEndpoint,
			decodeGetAllTodosRequest,
			encodeResponse,
			options...,
		),
	).Methods("GET")
}

type TodoEndpoints struct {
	CreateTodoEndpoint  endpoint.Endpoint
	UpdateTodoEndpoint  endpoint.Endpoint
	DeleteTodoEndpoint  endpoint.Endpoint
	GetTodoByIDEndpoint endpoint.Endpoint
	GetAllTodosEndpoint endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service. Useful in a server.
func MakeServerEndpoints(s todo.Service) TodoEndpoints {
	return TodoEndpoints{
		CreateTodoEndpoint:  MakeCreateTodoEndpoint(s),
		UpdateTodoEndpoint:  MakeUpdateTodoEndpoint(s),
		DeleteTodoEndpoint:  MakeDeleteTodoEndpoint(s),
		GetTodoByIDEndpoint: MakeGetTodoByIDEndpoint(s),
		GetAllTodosEndpoint: MakeGetAllTodosEndpoint(s),
	}
}

func MakeCreateTodoEndpoint(s todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(todo.CreateTodoRequest)
		response, err = s.CreateTodo(ctx, req)
		return
	}
}

func MakeUpdateTodoEndpoint(s todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(todo.UpdateTodoRequest)
		response, err = s.UpdateTodo(ctx, req)
		return
	}
}

func MakeDeleteTodoEndpoint(s todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(todo.DeleteTodoRequest)
		err = s.DeleteTodo(ctx, req)
		return
	}
}

func MakeGetTodoByIDEndpoint(s todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(todo.GetTodoByIDRequest)
		response, err = s.GetTodoByID(ctx, req)
		return
	}
}

func MakeGetAllTodosEndpoint(s todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (response interface{}, err error) {
		response, err = s.GetAllTodos(ctx)
		return
	}
}

func decodeCreateTodoRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req todo.CreateTodoRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, todo.Errorf(todo.EINVALID, "Failed to encode JSON body.")
	}

	return req, nil
}

func decodeUpdateTodoRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req todo.UpdateTodoRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, todo.Errorf(todo.EINVALID, "Failed to encode JSON body.")
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, todo.Errorf(todo.EINVALID, "Invalid value for parameter 'id'.")
	}

	req.ID, err = strconv.Atoi(id)
	if err != nil {
		return nil, todo.Errorf(todo.EINVALID, "Failed to convert '%s' to type integer.", id)
	}

	return req, nil
}

func decodeDeleteTodoRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req todo.DeleteTodoRequest

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, todo.Errorf(todo.EINVALID, "Invalid value for parameter 'id'.")
	}

	req.ID, err = strconv.Atoi(id)
	if err != nil {
		return nil, todo.Errorf(todo.EINVALID, "Failed to convert '%s' to type integer.", id)
	}

	return req, nil
}

func decodeGetTodoByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req todo.GetTodoByIDRequest

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, todo.Errorf(todo.EINVALID, "Invalid value for parameter 'id'.")
	}

	req.ID, err = strconv.Atoi(id)
	if err != nil {
		return nil, todo.Errorf(todo.EINVALID, "Failed to convert '%s' to type integer.", id)
	}

	return req, nil
}

func decodeGetAllTodosRequest(_ context.Context, _ *http.Request) (request interface{}, err error) {
	return nil, nil
}
