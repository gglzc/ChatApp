package user

import (
	"context"
	"strconv"
	"time"
	"fmt"

	"github.com/gglzc/StreamingWeb/util"
	"github.com/golang-jwt/jwt/v4"
)

const(
	secretKey = "fkwthewsorwld"
)

type service struct{
	Repository
	timeout	time.Duration
}

type MyJWTclaims struct{
	ID 			string		`json:"id"` 
	Username	string		`json:"username"`
	jwt.RegisteredClaims
}
func NewService(repository Repository)Service{
	return &service{
		repository,
		time.Duration(2)*time.Second,
	}
}
//
func (s *service) CreateUser (ctx context.Context ,req *CreateUserReq)(*CreateUserRes , error){
	ctx,cancel:=context.WithTimeout(ctx,s.timeout)
	defer cancel()
	//check username and email exist or not
	if err := s.CheckUsernameAndEmailExist(ctx, req.Username, req.Email); err != nil {
		return nil, err
	}
	//password 加密
	hashedPassword , err := util.HashPassword(req.Password)
	if err!=nil{
		return nil , err
	}
	
	u := &User{
		Username: req.Username,
		Email: 	  req.Email,
		Password: hashedPassword,
	}

	r , err:=s.Repository.CreateUser(ctx,u)
	if err!=nil{
		return nil,err
	}

	res:=&CreateUserRes{
		ID:			strconv.Itoa(int(r.ID)),
		Username: 	r.Username,
		Email: 		r.Email,
	}
	return res , nil	
}

func (s *service) Login(c context.Context,req *LoginUserReq) (*LoginUserRes , error){
	ctx , cancel := context.WithTimeout(c , s.timeout)
	defer cancel()

	u,err:=s.Repository.GetUserByEmail(ctx , req.Email)
	if err!=nil{
		return &LoginUserRes{},err
	}
	
	err = util.CheckPassword(req.Password,u.Password)
	if err!=nil{
		return &LoginUserRes{},err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512,MyJWTclaims{
		ID: strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
	})
	
	ss,err:=token.SignedString([]byte(secretKey))
	if err!=nil{
		return &LoginUserRes{},err
	}

	return &LoginUserRes{
		accessToken: ss,
		ID: strconv.Itoa(int(u.ID)),
		Username: u.Username,
	},nil
}

func (s *service) CheckUsernameAndEmailExist(ctx context.Context, username, email string) error {
    //check username is exist or not
    checkUsernameExist, err := s.Repository.CheckUsernameExist(ctx, username)
    if err != nil {
        return err
    }
    if !checkUsernameExist {
        return fmt.Errorf("the username is already exists")
    }

    //check email is exist or not
    checkEmailExist, err := s.Repository.CheckEmailExist(ctx, email)
    if err != nil {
        return err
    }
    if !checkEmailExist {
        return fmt.Errorf("the email is already exists")
    }

    return nil
}