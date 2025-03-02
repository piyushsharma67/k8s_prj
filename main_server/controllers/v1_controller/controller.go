package v1_controller

import (
	"main_server/controllers/common"
	"main_server/proto"
	"main_server/services"
)


type V1Controller struct{
	*common.ControllerStruct
}

func InitialiseV1Controller(service *services.ServiceStruct, auth proto.AuthServiceClient)*V1Controller{
	return &V1Controller{
		ControllerStruct: &common.ControllerStruct{
			AuthService: auth,
			Service: service,
		},
	}
}