package routes

import (
	"fmt"
	"main_server/controllers"
	"main_server/proto"
	v1 "main_server/routes/v1"
	"main_server/services"

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

func InitRoutes(service *services.ServiceStruct, auth proto.AuthServiceClient) *mux.Router {
	r := mux.NewRouter()
	r.Use(LoggingMiddleware)

	r.HandleFunc("/", controllers.Health).Methods("GET")

	v1Route := r.PathPrefix("/v1").Subrouter()
	v1.RegisterRoutes(v1Route, service, auth)

	return r
}
