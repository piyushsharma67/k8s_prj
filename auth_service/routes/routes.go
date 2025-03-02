package routes

import (
	"auth_service/controllers/http_controller"
	"auth_service/services"
	"fmt"

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

func InitRoutes(service *services.ServiceStruct) *mux.Router {
	r := mux.NewRouter()
	r.Use(LoggingMiddleware)

	contoller := http_controller.HTTPController{}
	c := contoller.NewHTTPController(service)

	r.HandleFunc("/", c.Health)
	r.HandleFunc("/signup", c.SignupHttp).Methods("POST")
	r.HandleFunc("/login", c.LoginHttp).Methods("GET")
	r.HandleFunc("/save_fc_token", Protected(c.SaveUserFcmHttp)).Methods("POST")

	return r
}
