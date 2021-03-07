package logmw

import (
	"context"
	"github.com/go-kit/kit/log"
	"time"
	"todo"
)

func NewTodoLoggingMiddleware(logger log.Logger) todo.Middleware {
	return func(next todo.Service) todo.Service {
		return &todoLoggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type todoLoggingMiddleware struct {
	next   todo.Service
	logger log.Logger
}

func (mw todoLoggingMiddleware) CreateTodo(ctx context.Context, request todo.CreateTodoRequest) (t *todo.Todo, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "CreateTodo",
			"value", request.Value,
			"complete", request.Complete,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return mw.next.CreateTodo(ctx, request)
}

func (mw todoLoggingMiddleware) UpdateTodo(ctx context.Context, request todo.UpdateTodoRequest) (t *todo.Todo, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "UpdateTodo",
			"id", request.ID,
			"value", request.Value,
			"complete", request.Complete,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return mw.next.UpdateTodo(ctx, request)
}

func (mw todoLoggingMiddleware) DeleteTodo(ctx context.Context, request todo.DeleteTodoRequest) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "DeleteTodo",
			"id", request.ID,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return mw.next.DeleteTodo(ctx, request)
}

func (mw todoLoggingMiddleware) GetTodoByID(ctx context.Context, request todo.GetTodoByIDRequest) (t *todo.Todo, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetTodoByID",
			"id", request.ID,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return mw.next.GetTodoByID(ctx, request)
}

func (mw todoLoggingMiddleware) GetAllTodos(ctx context.Context) (t []*todo.Todo, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetAllTodos",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())

	return mw.next.GetAllTodos(ctx)
}
