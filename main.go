package main

import (
	"log"

	_"github.com/gglzc/StreamingWeb/docs"
	"github.com/gglzc/StreamingWeb/db"
	"github.com/gglzc/StreamingWeb/db/redis"
	"github.com/gglzc/StreamingWeb/internal/user"
	"github.com/gglzc/StreamingWeb/internal/ws"
	"github.com/gglzc/StreamingWeb/router"
)
// @title OMA CHAT API
// @version 2.0
// @description Gin swagger
// @contact.name gglzc
// @host localhost:8085
// schemes http
func main(){
	dbConn,err:=db.NewDatebase()
	if err!=nil{
		log.Fatalf("can't connect db cause : %s" ,err)
	}
	
	redisdbConn  := redis.NewCache()
	
	userRep := user.NewRepository(dbConn.GetDB(),redisdbConn.GetChache())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub:=ws.NewHub()
	wshandler :=ws.NewHandler(hub)
	go hub.Run()
	
	router.InitRouter(userHandler , wshandler)
	router.Start("0.0.0.0:8085")
}