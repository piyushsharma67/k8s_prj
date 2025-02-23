package services

import (
	"auth_service/repository"
)

type ServiceStruct struct {
	Repository *repository.Repositories
}

func (s *ServiceStruct) InitialiseService(r *repository.Repositories) *ServiceStruct {
	s.Repository = r
	return s
}
