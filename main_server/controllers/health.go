package controllers

import (
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("i am working fine! from main-server!!"))
}
