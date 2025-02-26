package services

import (
	"main_server/proto"
	"main_server/repository"
)

type ServiceStruct struct {
	proto.UnimplementedAuthServiceServer
	Repository *repository.Repositories
}

func (s *ServiceStruct) InitialiseService(r *repository.Repositories) *ServiceStruct {
	s.Repository = r
	return s
}
