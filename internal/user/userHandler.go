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
// CreateUser creates a new user account.
// @Summary 申請會員
// @Tags Signup
// @Router /signup [post]
// @Accept json
// @Produce json
// @Param username body string true "Username"
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {string} user.CreateUserRes
// @Failure 400 {object} string "error"
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
// Login logs in a user and returns an access token.
// @Summary 登入會員
// @Tags Login
// @Router /login [post]
// @Accept json
// @Produce json
// @Param email formData string true "User's email address"
// @Param password formData string true "User's password"
// @Success 200 {object} user.LoginUserRes
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
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
// LogOut logs out the current user.
// @Summary 登出會員
// @Tags Logout
// @Router /logout [GET]
// @Success 200 {object} string "Logout successful"
func (h *Handler) LogOut(c *gin.Context){
		c.SetCookie("jwt","",-1,"","",false,true)
		c.JSON(http.StatusOK , gin.H{"message" : "log out successful"})
}
