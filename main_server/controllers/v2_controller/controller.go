package v2_controller

import (
	"main_server/controllers/common"
	"main_server/proto/auth"
	"main_server/services"
)

type V1Controller struct {
	*common.ControllerStruct
}

func InitialiseV1Controller(service *services.ServiceStruct, auth auth.AuthServiceClient) *V1Controller {
	return &V1Controller{
		ControllerStruct: &common.ControllerStruct{
			AuthService: auth,
			Service:     service,
		},
	}
}
