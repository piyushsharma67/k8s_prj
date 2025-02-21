package routes

import (
	"fmt"
	"k8s_project/auth_service/controllers"
	"k8s_project/auth_service/services"
	"net/http"

	"time"

	"github.com/gorilla/mux"
)

var P string

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func InitRoutes(service *services.ServiceStruct)*mux.Router{
	r:=mux.NewRouter()
	r.Use(LoggingMiddleware)

	contoller:=controllers.ControllerStruct{}
	c:=contoller.InitialiseController(service)

	r.HandleFunc("/",c.Health)
	r.HandleFunc("/v1/signup",c.SignupUser).Methods("POST")
	r.HandleFunc("/v1/save_fc_token",Protected(c.SaveUserFcmToken)).Methods("POST")
	
	return r
}