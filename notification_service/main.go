package main

import (
	"log"
	"notification_service/cmd"
)

func main(){
	if err:=cmd.Execute();err!=nil{
		log.Fatal(err)
	}
}