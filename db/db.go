package db

import (
	"database/sql"
  	_"github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatebase()(*Database , error ){
	db,err:=sql.Open("postgres","postgresql://root:test@localhost:5430/go-chat?sslmode=disable")
	if err!=nil{
		return nil,err
	}
	return &Database{db: db},err
}

func(d *Database)Close(){
	d.db.Close()
}

func(d *Database)GetDB()*sql.DB{
	return d.db
}