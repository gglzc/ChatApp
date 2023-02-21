package user

import (
	"context"
	"database/sql"
	"time"

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
	Set(ctx context.Context,key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Exists(ctx context.Context,key ...string) (*redis.IntCmd)
	SAdd(ctx context.Context,key string,member ...interface{})*redis.IntCmd
	SIsMember(ctx context.Context,key string,member interface{})*redis.BoolCmd
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
	var existUsername string
	query:="SELECT username FROM public.user WHERE username = $1"
	err:=r.db.QueryRowContext(ctx , query , username).Scan(&existUsername)
	if err!=nil && err != sql.ErrNoRows{
	//mean user doesn't exist
		return false,err
	}
	//mean username exist
	return existUsername!="",nil
}

func (r *repository)CheckEmailExist(ctx context.Context , email string)(bool,error){
	var existEmail string
	query:="SELECT email FROM public.user WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&existEmail)
	 //check if there is error first 
	if err != nil && err!=sql.ErrNoRows{
        return false, err
    }
	//mean email  exist 
	return existEmail!="",nil
}


func (r *repository)CheckEmailByCache(ctx context.Context, email string) (bool,error){
	exist,err:=r.redisdb.SIsMember(ctx,"email",email).Result()
	if err!=nil{
		return false,err
	}
	return exist,nil
}

func (r *repository)CheckUsernameByCache(ctx context.Context ,username string) (bool,error){
	exist,err:=r.redisdb.SIsMember(ctx,"username",username).Result()
	if err!=nil{
		return false,err
	}
	return exist,nil
}

func (r *repository)CreateUserByCache(ctx context.Context, user *User) error{
	//username
	_,err:=r.redisdb.SAdd(ctx,"username",user.Username).Result()
	if err!=nil{
		return err
	}
	//email
	_, err = r.redisdb.SAdd(ctx, "email", user.Email).Result()
    if err != nil {
        return err
    }

    return nil
}