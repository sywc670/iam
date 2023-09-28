package policy

import (
	srvv1 "github.com/marmotedu/iam/internal/will-apiserver/service/v1"
	"github.com/marmotedu/iam/internal/will-apiserver/store"
)

type PolicyController struct {
	srv srvv1.Service
}

func NewPolicyController(store store.Factory) *PolicyController {
	return &PolicyController{
		srv: srvv1.NewService(store),
	}
}
