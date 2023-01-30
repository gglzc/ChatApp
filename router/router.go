package router

import (
	"time"
	
	"github.com/gglzc/StreamingWeb/internal/user"
	"github.com/gglzc/StreamingWeb/internal/ws"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler , wsHandler *ws.Handler){
	r = gin.Default()
	
	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
		  return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	  }))
	
	r.POST("/signup" ,userHandler.CreateUser)
	r.POST("/login",userHandler.Login)
	r.GET("/logout",userHandler.LogOut)

	r.POST("/ws/createroom" , wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomId",wsHandler.JoinRoom)
	r.GET("/ws/getRooms" , wsHandler.GetRooms)
	r.GET("/ws/getClients/:roomId" , wsHandler.GetRoomsClients)

}

func Start(addr string) error{
	return  r.Run(addr)
}