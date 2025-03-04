package routes

import (
	"main_server/controllers"
	"main_server/proto/auth"
	v1 "main_server/routes/v1"
	"main_server/services"
	"main_server/utils"

	"github.com/gorilla/mux"
)

var P string

func InitRoutes(service *services.ServiceStruct, auth auth.AuthServiceClient) *mux.Router {
	r := mux.NewRouter()
	r.Use(utils.LoggingMiddleware)

	r.HandleFunc("/", controllers.Health).Methods("GET")

	v1Route := r.PathPrefix("/v1").Subrouter()
	v1.RegisterRoutes(v1Route, service, auth)

	return r
}
