package redis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"os"
)

var rdb *redis.Client

func Init() {
	var env, addr string
	env = os.Getenv("GIN_MODE")
	if env == gin.ReleaseMode {
		addr = "redis-tiny-vote:6379"
	} else {
		addr = "localhost:6379"
	}
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		panic("Fail to connect to the redis, error is " + err.Error())
	}
}
