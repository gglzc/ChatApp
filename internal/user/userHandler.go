package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
			Service: s,
	}
}

func (h *Handler)CreateUser(c *gin.Context){
	var req CreateUserReq
	if err:=c.ShouldBind(&req) ; err!=nil{
		c.JSON(http.StatusBadRequest , gin.H{
			"error":err.Error()})
		return
	}
	//res
	res , err := h.Service.CreateUser(c.Request.Context() , &req)
	if err!=nil{
		c.JSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
		return 
	}
	c.JSON(http.StatusOK , res)
}
// For login 
func (h *Handler)Login(c *gin.Context){
	var user LoginUserReq
	
	if err:=c.ShouldBindJSON(&user);err!=nil{
		c.JSON(http.StatusBadRequest , gin.H{"error":err.Error()})
		return
	}

	u,err:=h.Service.Login(c.Request.Context(),&user)
	if err!=nil{
		c.JSON(http.StatusInternalServerError , gin.H{"error":err.Error()})
		return
	}

	c.SetCookie("jwt" , u.accessToken , 3600 ,"/","localhost",false,true)
	c.JSON(http.StatusOK , u)
}
//for Service
func (h *Handler) LogOut(c *gin.Context){
		c.SetCookie("jwt","",-1,"","",false,true)
		c.JSON(http.StatusOK , gin.H{"message" : "log out successful"})
	}