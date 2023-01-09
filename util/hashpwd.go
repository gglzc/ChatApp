package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string)(string , error){
	hashPassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		return "",fmt.Errorf("fail to hashed password : %w" , err)
	}

	return string(hashPassword),nil
}

func CheckPassword(password string , hashPassword string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashPassword),[]byte(password))
}