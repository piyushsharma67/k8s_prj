package main

import (
	"k8s_project/auth_service/cmd"
	"log"
)

func main(){
	if err:=cmd.Execute();err!=nil{
		log.Fatal(err)
	}
}