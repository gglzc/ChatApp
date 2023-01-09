package db

import "github.com/go-redis/redis/v9"

type Redisdb struct{
	rd *redis.Client
}

func  NewCache() (*Redisdb ){
	rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
	return &Redisdb{rd: rdb,}
}

func (nc *Redisdb) Close(){
	 nc.rd.Close()
}

func (nc *Redisdb) GetChache(){

}