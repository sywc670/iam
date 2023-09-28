package v1

import "github.com/marmotedu/iam/internal/will-apiserver/store"

type Service interface {
	Users() UserSrv
	Secrets() SecretSrv
	Policies() PolicySrv
}

type service struct {
	store store.Factory
}

func (s *service) Users() UserSrv {
	return newUsers(s)
}

func (s *service) Secrets() SecretSrv {
	return newSecrets(s)
}

func (s *service) Policies() PolicySrv {
	return newPolicies(s)
}

func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}
