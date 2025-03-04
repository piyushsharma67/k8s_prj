package common

import (
	"auth_service/proto/auth"
	"auth_service/services"
)


type ControllerStruct struct{
	auth.UnimplementedAuthServiceServer
	Service *services.ServiceStruct
}

func (c *ControllerStruct)InitialiseController(s *services.ServiceStruct)*ControllerStruct{
	c.Service=s

	return c
}