package model

import (
	"context"
	"github.com/go-redis/redis/v9"
)

type Sensor struct {
	ID           string `json:"id"`
	IATA         string `json:"iata"`
	MEASURETYPE  string `json:"measuretype"`
	MEASUREVALUE string `json:"measurevalue"`
	MEASUREDATE  string `json:"measuredate"`
}

func (s *Sensor) GetSensor(client *redis.Client) error {
	return client.Get(context.Background(), "sensor:"+s.ID).Scan(s)
}

func GetSensors(client *redis.Client) ([]Sensor, error) {
	var sensors []Sensor
	ctx := context.Background()
	keys, err := client.Keys(ctx, "sensor:*").Result()
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		var s Sensor
		err = client.Get(ctx, key).Scan(&s)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, s)
	}
	return sensors, nil
}
