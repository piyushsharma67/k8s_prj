package main

import (
	"auth_service/cmd"
	"log"
)

func main(){
	if err:=cmd.Execute();err!=nil{
		log.Fatal(err)
	}
}