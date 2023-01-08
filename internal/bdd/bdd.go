package bdd

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func AddToSortedSet(name string, score int64, value string) int64 {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	r, err := redis.Int64(conn.Do("ZADD", name, score, value))
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func GetValuesBetween2Index(name string, firstIndex, secondIndex int64) []string {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	r, err := redis.Values(conn.Do("ZRANGE", name, firstIndex, secondIndex))
	if err != nil {
		log.Fatal(err)
	}

	return scanSlice(r)
}

func GetValuesBetween2Score(name string, firstScore, secondScore int64) []string {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	r, err := redis.Values(conn.Do("ZRANGEBYSCORE", name, firstScore, secondScore))
	if err != nil {
		log.Fatal(err)
	}

	return scanSlice(r)
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

func scanSlice(values []interface{}) []string {
	var results []string
	var err error

	for len(values) > 0 {
		var value string

		values, err = redis.Scan(values, &value)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, value)
	}
	return results
}
