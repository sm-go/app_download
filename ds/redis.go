package ds

import (
	"app-download/config"
	"fmt"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func ConnectRedis() *redis.Client {
	addr := fmt.Sprintf("%s:%s", config.RedisHost.Host, config.RedisHost.Port)
	pass := config.RedisHost.Pass
	db, err := strconv.ParseInt(config.RedisHost.DB, 10, 64)

	if err != nil {
		log.Fatalln("error on connection redis", err.Error())
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       int(db),
	})
	RDB = rdb
	return rdb
}
