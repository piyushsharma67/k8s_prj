package controllers

import "k8s_project/auth_service/services"


type ControllerStruct struct{
	service *services.ServiceStruct
}

func (c *ControllerStruct)InitialiseController(s *services.ServiceStruct)*ControllerStruct{
	c.service=s

	return c
}