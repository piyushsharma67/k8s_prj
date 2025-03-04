package http_controller

import "net/http"


func (c *HTTPController)Health(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("i am working fine!"))
}