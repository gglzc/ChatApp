package user

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis/v9"
)
type DBTX interface{
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type ChacheTx interface{
    Get(ctx context.Context, key string) (*redis.StringCmd)
}

type repository struct{
	db DBTX
	redisdb	ChacheTx
}

func NewRepository(db DBTX ,redisdb ChacheTx) Repository{
	return &repository{
		db:      db,
		redisdb: redisdb,
	}
}

func (r *repository) CreateUser(ctx context.Context,user *User)(*User , error ){
	var lastInsertID int

	query := "INSERT INTO public.user(username, password, email) VALUES ($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx , query, user.Username,user.Password,user.Email).Scan(&lastInsertID)
	if err!=nil{
		return &User{},err
	}

	user.ID=int64(lastInsertID)
	return user,nil
}

func (r *repository) GetUserByEmail(ctx context.Context , email string) (*User , error){
	u:=User{}
	
	query:="SELECT id, email, username, password FROM public.user WHERE email = $1"
	err:=r.db.QueryRowContext(ctx,query,email).Scan(&u.ID,&u.Email,&u.Username,&u.Password)
	if err!=nil{
		return &User{},nil
	}
	return &u,nil
}

func (r *repository)CheckUsernameExist(ctx context.Context , username string) (bool,error){
	query:="SELECT username FROM public.user WHERE username = $1"
	rows,err:=r.db.QueryContext(ctx , query , username)
	 
	 //check if there is error first
	 if err != nil {
        return false, err
    }
    defer rows.Close()
	//it's mean that the username is already exist
	if rows.Next() {
        return false, nil
    }
	return true,nil
}

func (r *repository)CheckEmailExist(ctx context.Context , email string)(bool,error){
	query:="SELECT email FROM public.user WHERE email = $1"
	rows, err := r.db.QueryContext(ctx, query, email)
	 //check if there is error first 
	 if err != nil {
        return false, err
    }
    defer rows.Close()
	//mean that the email does  exist
	if rows.Next() {
        return false, nil
    }
	//mean that the email doesn't  exist 
	return true,nil
}


func (r *repository)CheckEmailByCache(ctx context.Context, email string) (bool,error){
	key:=r.redisdb.Get(ctx , email)
	emailExist , err := key.Bool()
	if err!=nil{
		return false , err
	}
	return emailExist ,err
}

func (r *repository)CheckUsernameByCache(ctx context.Context ,username string) (bool,error){
	key:=r.redisdb.Get(ctx , username )
	usernameExist , err := key.Bool()
	if err!=nil{
		return false , err
	}
	return usernameExist,err
}