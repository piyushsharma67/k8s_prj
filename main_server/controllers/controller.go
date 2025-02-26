package controllers

import (
	"main_server/proto"
	"main_server/services"
)

type ControllerStruct struct{
	service *services.ServiceStruct
	authService proto.AuthServiceClient
}

func (c *ControllerStruct)InitialiseController(s *services.ServiceStruct,p proto.AuthServiceClient)*ControllerStruct{
	c.service=s
	c.authService=p
	return c
}