package user

import "context"

type User struct{
	ID 			int64		`json:"id" db:"id"` 
	Username	string		`json:"username" db:"username"`
	Email		string		`json:"email" db:"email"`
	Password	string		`json:"password" db:"password"`
}

type CreateUserReq struct{
	Username	string		`json:"username" db:"username"`
	Email		string		`json:"email" db:"email"`
	Password	string		`json:"password" db:"password"`
}

type CreateUserRes struct{
	ID 			string		`json:"id" db:"id"` 
	Username	string		`json:"username" db:"username"`
	Email		string		`json:"email" db:"email"`
}

type LoginUserReq struct{
	Email 		string		`json:"email" db:"email"`
	Password	string		`json:"password" db:"password"`
}

type LoginUserRes struct{
	accessToken	string		
	ID 			string		`json:"id" db:"id"` 
	Username	string		`json:"username" db:"username"`
}

type Repository interface{
	CreateUser(ctx context.Context,user *User)(*User , error )
	GetUserByEmail(ctx context.Context , email string) (*User , error)
	CheckUsernameExist(ctx context.Context , username string) (bool,error)
	CheckEmailExist(ctx context.Context , email string) (bool,error)
	
	CheckEmailByCache(ctx context.Context ,key string) (bool, error)
	CheckUsernameByCache(ctx context.Context ,key string) (bool,error)
	CreateUserByCache(ctx context.Context, user *User) error
}


type Service interface{
	CreateUser(ctx context.Context ,req *CreateUserReq)(*CreateUserRes , error)
	Login(ctx context.Context,req *LoginUserReq) (*LoginUserRes , error)
}