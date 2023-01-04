package config

import (
	"github.com/tkanos/gonfig"
)

type Captor struct {
	CLIENT_ID string
}

type Configuration struct {
	ADDRESS     string
	PORT        string
	QOS         byte
	DELAY       int
	PRESSURE    Captor
	TEMPERATURE Captor
	WIND        Captor
}

func GetConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("../internal/config/config.json", &configuration)
	return configuration
}
