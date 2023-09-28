package user

import (
	srvv1 "github.com/marmotedu/iam/internal/will-apiserver/service/v1"
	"github.com/marmotedu/iam/internal/will-apiserver/store"
)

type UserController struct {
	srv srvv1.Service
}

func NewUserController(store store.Factory) *UserController {
	return &UserController{
		srv: srvv1.NewService(store),
	}
}
