package service

import (
	"fmt"

	"github.com/go-redis/redis"
)

// ConnectToRedis Exported
func ConnectToRedis() *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return conn
}

// AddToRedis Exported
func AddToRedis(key, value string) {
	conn := ConnectToRedis()
	if err := conn.Set(key, value, 0).Err(); err != nil {
		panic(err)
	}
}

// SearchInRedis Exported
func SearchInRedis(value string) {
	conn := ConnectToRedis()
	if key := conn.Get(value).Val(); key != "" {
		fmt.Println(conn.Get(value).Val())
	} else {
		fmt.Println("Indice Não Encontrado!")
	}
}
