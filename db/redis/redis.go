package db

import (
	"github.com/go-redis/redis/v9"
)

type Redisdb struct {
	redisdb *redis.Client
}

// Get implements user.ChacheTx

func NewCache() *Redisdb {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &Redisdb{redisdb: rdb}
}

func (rdb *Redisdb) Close() {
	rdb.redisdb.Close()
}

func (rdb *Redisdb) GetChache() *Redisdb {
	return rdb
}
