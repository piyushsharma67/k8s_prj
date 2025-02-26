package controllers

import "net/http"

func (c *ControllerStruct) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("i am working fine! from main-server!!"))
}
