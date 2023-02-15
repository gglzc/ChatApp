package router

import (
	"time"
	
	"github.com/gglzc/StreamingWeb/internal/user"
	"github.com/gglzc/StreamingWeb/internal/ws"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/zap"
	"go.uber.org/zap"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/gglzc/StreamingWeb/docs"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler , wsHandler *ws.Handler){
	r = gin.Default()
	//zaplogger
	logger , _:= zap.NewProduction()
	r.Use(ginzap.Ginzap(logger,time.RFC3339,true))
	//Cors
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
	//swagger
	url := ginSwagger.URL("http://localhost:8085/swagger/doc.json") // The url pointing to API definition
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//router
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