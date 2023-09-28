package secret

import (
	srvv1 "github.com/marmotedu/iam/internal/will-apiserver/service/v1"
	"github.com/marmotedu/iam/internal/will-apiserver/store"
)

type SecretController struct {
	srv srvv1.Service
}

func NewSecretController(store store.Factory) *SecretController {
	return &SecretController{
		srv: srvv1.NewService(store),
	}
}
