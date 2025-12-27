package controller

import (
	"github.com/jobpay/todo/internal/presentation/controller/todo"
)

type Controllers struct {
	Todo *todo.Controllers
}

func NewControllers(todo *todo.Controllers) *Controllers {
	return &Controllers{
		Todo: todo,
	}
}
