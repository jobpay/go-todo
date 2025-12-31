package controller

import (
	"github.com/jobpay/todo/internal/presentation/controller/tag"
	"github.com/jobpay/todo/internal/presentation/controller/todo"
)

type Controllers struct {
	Todo *todo.Controllers
	Tag  *tag.Controllers
}

func NewControllers(todo *todo.Controllers, tag *tag.Controllers) *Controllers {
	return &Controllers{
		Todo: todo,
		Tag:  tag,
	}
}
