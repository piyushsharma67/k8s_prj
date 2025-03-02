package v1

import (
	"main_server/controllers/v1_controllers"
	"main_server/proto"
	"main_server/services"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, service *services.ServiceStruct, auth proto.AuthServiceClient) {

	c:=v1_controller.InitialiseV1Controller(service,auth)
	
	r.HandleFunc("/signup", c.SignupUser).Methods("POST")
	
}
