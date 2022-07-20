package redisclient

import (
	"fmt"
	"sync"

	"github.com/ajikamaludin/go-fiber-rest/app/configs"
	"github.com/go-redis/redis/v8"
)

var lock = &sync.Mutex{}
var rdb *redis.Client

func GetInstance() *redis.Client {
	fmt.Println("redis client", rdb)
	if rdb == nil {
		configs := configs.GetInstance()
		addr := fmt.Sprintf("%s:%s", configs.Redisconfig.Host, configs.Redisconfig.Port)
		lock.Lock()
		rdb = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: configs.Redisconfig.Password, // no password set
			DB:       0,                            // use default DB
		})
		lock.Unlock()
	}
	return rdb
}
