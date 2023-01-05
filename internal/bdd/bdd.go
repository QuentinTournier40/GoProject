package bdd

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func SetValue(key string, value string) {
	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	_, err = redis.String(conn.Do("SET", key, value))
}

func GetValue(key string) string {
	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	r, err := redis.String(conn.Do("GET", key))
	return r
}

func GetAllKeyRegex(expression string) []string {
	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	r, err := redis.Strings(conn.Do("KEYS", expression))
	return r
}
