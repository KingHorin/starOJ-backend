package config

import (
	"github.com/go-redis/redis"
)

var rd *redis.Client

func GetRD() *redis.Client {
	return rd
}

func init() {

	rd = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // redis一共16个库，指定其中一个库即可
	})
	_, err := rd.Ping().Result()
	if err != nil {
		panic("redis连接失败, error = " + err.Error())
	}

	rd.FlushDB() //清空所有缓存

}
