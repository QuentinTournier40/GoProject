package main

import (
	"context"
	"github.com/go-redis/redis/v9"
)

type sensor struct {
	ID           string `json:"id"`
	IATA         string `json:"iata"`
	MEASURETYPE  string `json:"measuretype"`
	MEASUREVALUE string `json:"measurevalue"`
	MEASUREDATE  string `json:"measuredate"`
}

func (s *sensor) getSensor(client *redis.Client) error {
	return client.Get(context.Background(), "sensor:"+s.ID).Scan(s)
}

func getSensors(client *redis.Client) ([]sensor, error) {
	var sensors []sensor
	ctx := context.Background()
	keys, err := client.Keys(ctx, "sensor:*").Result()
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		var s sensor
		err = client.Get(ctx, key).Scan(&s)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, s)
	}
	return sensors, nil
}
