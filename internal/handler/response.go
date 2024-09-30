package handler

import (
	"net/http"

	"intercom/internal/models/response"
)

type HandlerFunc[T any] func(w http.ResponseWriter, r *http.Request) Response[T]

type Response[T any] struct {
	Data          *T
	Error         *response.Error
	InternalError *response.Error
	Status        int
}

type ResponseBuilder[T any] struct {
	response Response[T]
}

func NewResponseBuilder[T any]() *ResponseBuilder[T] {
	return &ResponseBuilder[T]{}
}

func (b *ResponseBuilder[T]) SetData(data *T) *ResponseBuilder[T] {
	b.response.Data = data
	return b
}

func (b *ResponseBuilder[T]) SetInternalError(err error) *ResponseBuilder[T] {
	if err != nil {
		b.response.InternalError = &response.Error{
			Message: err.Error(),
		}
	}
	return b
}

func (b *ResponseBuilder[T]) SetError(err error) *ResponseBuilder[T] {
	if err != nil {
		b.response.Error = &response.Error{
			Message: err.Error(),
		}
	}
	return b
}

func (b *ResponseBuilder[T]) SetStatus(status int) *ResponseBuilder[T] {
	b.response.Status = status
	return b
}

func (b *ResponseBuilder[T]) Build() Response[T] {
	return b.response
}
