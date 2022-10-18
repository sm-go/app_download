package ds

import (
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type DataSource struct {
	DB  *gorm.DB
	RDB *redis.Client
}

func NewDataSource() *DataSource {
	db, err := ConnectMysql()
	if err != nil {
		log.Fatal("Can't connect mysql")
	}

	rdb := ConnectRedis()

	return &DataSource{
		DB:  db,
		RDB: rdb,
	}
}
