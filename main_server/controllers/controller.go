package controllers

import (
	"main_server/proto"
	"main_server/services"
)




type ControllerStruct struct{
	service *services.ServiceStruct
	authService proto.AuthServiceClient
}

func (c *ControllerStruct)InitialiseController(s *services.ServiceStruct)*ControllerStruct{
	c.service=s

	return c
}