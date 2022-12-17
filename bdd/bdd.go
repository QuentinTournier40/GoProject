package bdd

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

func SetValue(key string, value string) {
	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	r, err := redis.String(conn.Do("SET", key, value))
	fmt.Println(r)
}
