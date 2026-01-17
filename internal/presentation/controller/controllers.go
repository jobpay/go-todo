package controller

import (
	"github.com/jobpay/todo/internal/presentation/controller/auth"
	"github.com/jobpay/todo/internal/presentation/controller/tag"
	"github.com/jobpay/todo/internal/presentation/controller/todo"
)

type Controllers struct {
	Todo *todo.Controllers
	Tag  *tag.Controllers
	Auth *auth.Controllers
}

func NewControllers(todo *todo.Controllers, tag *tag.Controllers, auth *auth.Controllers) *Controllers {
	return &Controllers{
		Todo: todo,
		Tag:  tag,
		Auth: auth,
	}
}
