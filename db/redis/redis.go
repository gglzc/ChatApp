package redis

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

var(
	rdb	*Redisdb
	once sync.Once
)
type Redisdb struct {
	redisdb *redis.Client
}

func init(){
// 設置Viper
	viper.SetConfigName("config") // 文件名稱
	viper.SetConfigType("yaml") // 文件類型
	viper.AddConfigPath("./config") // 搜索的路徑，此處為當前目錄

// 讀取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fail to load config:%s", err))
	}
}


func NewCache() *Redisdb {
	once.Do(func ()  {
		redisAddr := viper.GetString("redis.address")
		redisPassword := viper.GetString("redis.password")
		redisDB := viper.GetInt("redis.db")
		
		client := redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: redisPassword, // no password set
			DB:       redisDB,  // use default DB
		})
		rdb = &Redisdb{redisdb: client}
	})
	return rdb
}

func (rdb *Redisdb) Close() {
	rdb.redisdb.Close()
}

func (rdb *Redisdb) GetChache() *redis.Client {

	return rdb.redisdb
}
