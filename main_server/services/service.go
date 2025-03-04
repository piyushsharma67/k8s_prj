package services

import (
	"main_server/proto/auth"
	"main_server/repository"
)

type ServiceStruct struct {
	auth.UnimplementedAuthServiceServer
	Repository *repository.Repositories
}

func (s *ServiceStruct) InitialiseService(r *repository.Repositories) *ServiceStruct {
	s.Repository = r
	return s
}
