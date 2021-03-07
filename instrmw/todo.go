package instrmw

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"time"
	"todo"
)

func NewTodoInstrumentingMiddleware(
	requestCount metrics.Counter,
	errorCount metrics.Counter,
	requestDuration metrics.Histogram,
) todo.Middleware {
	return func(next todo.Service) todo.Service {
		return todoInstrumentingMiddleware{
			requestCount:    requestCount,
			errorCount:      errorCount,
			requestDuration: requestDuration,
			service:         next,
		}
	}
}

type todoInstrumentingMiddleware struct {
	requestCount    metrics.Counter
	errorCount      metrics.Counter
	requestDuration metrics.Histogram
	service         todo.Service
}

func (mw todoInstrumentingMiddleware) CreateTodo(ctx context.Context, request todo.CreateTodoRequest) (t *todo.Todo, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CreateTodo", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	t, err = mw.service.CreateTodo(ctx, request)
	return
}

func (mw todoInstrumentingMiddleware) UpdateTodo(ctx context.Context, request todo.UpdateTodoRequest) (t *todo.Todo, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "UpdateTodo", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	t, err = mw.service.UpdateTodo(ctx, request)
	return
}

func (mw todoInstrumentingMiddleware) DeleteTodo(ctx context.Context, request todo.DeleteTodoRequest) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "DeleteTodo", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	err = mw.service.DeleteTodo(ctx, request)
	return
}

func (mw todoInstrumentingMiddleware) GetTodoByID(ctx context.Context, request todo.GetTodoByIDRequest) (t *todo.Todo, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetTodoByID", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	t, err = mw.service.GetTodoByID(ctx, request)
	return
}

func (mw todoInstrumentingMiddleware) GetAllTodos(ctx context.Context) (t []*todo.Todo, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAllTodos", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestDuration.With(lvs...).Observe(time.Since(begin).Seconds())
		if err != nil {
			mw.errorCount.With(lvs...).Add(1)
		}
	}(time.Now())
	t, err = mw.service.GetAllTodos(ctx)
	return
}
