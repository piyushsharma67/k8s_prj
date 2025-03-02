package http_controller

import (
	"auth_service/controllers/common"
	"auth_service/services"
)


type HTTPController struct {
    *common.ControllerStruct
}

func (c *HTTPController)NewHTTPController(s *services.ServiceStruct) *HTTPController {
    c.ControllerStruct = &common.ControllerStruct{}
    c.ControllerStruct.InitialiseController(s)

    return c
}