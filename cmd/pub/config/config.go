package config

import (
	"fmt"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	ADDRESS   string
	PORT      string
	QOS       byte
	CLIENT_ID string
	DELAY     int
}

func GetConfig(file string) Configuration {
	configuration := Configuration{}

	fileName := fmt.Sprintf("./cmd/pub/config/%s_config.json", file)
	gonfig.GetConf(fileName, &configuration)
	return configuration
}
