package Publishers

import (
	"fmt"
	"goproject/cmd/PubSubMethods"
	"goproject/config"
	"math/rand"
	"time"
)

func RunPublisher(captorFullNameUpperCase, captorId string, minValue, maxValue float64) {
	// GENERATE RANDOM SEED
	rand.Seed(time.Now().UnixNano())

	configuration := config.GetConfig()
	clientId := ""
	switch captorId {
	case "1":
		clientId = configuration.PRESSURE.CLIENT_ID

	case "2":
		clientId = configuration.TEMPERATURE.CLIENT_ID

	case "3":
		clientId = configuration.WIND.CLIENT_ID
	}

	topic := "capteurs"
	client := PubSubMethods.Connect(configuration.ADDRESS+":"+configuration.PORT, clientId, configuration.DELAY)

	mapIata := config.CODE_IATA

	var tabValue []float64

	// SET FIRST RANDOM VALUE
	for range mapIata {
		tabValue = append(tabValue, minValue+rand.Float64()*(maxValue-minValue))
	}

	for {
		for key, value := range mapIata {
			tabValue[key] = generateCoherenteValue(tabValue[key])
			now := time.Now()
			msg := captorId + " " + value + " " + captorFullNameUpperCase + " " + fmt.Sprintf("%.1f", tabValue[key]) + " " + now.Format("2006-01-02-15-04-05")
			client.Publish(topic, configuration.QOS, false, msg)
			fmt.Println(msg)
		}
		time.Sleep(time.Duration(configuration.DELAY) * time.Second)
	}
}

func generateCoherenteValue(value float64) float64 {
	randomInterval := rand.Float64() * 3
	val := 0.0
	if rand.Float64() < 0.5 {
		val = value - randomInterval
	} else {
		val = value + randomInterval
	}
	return val
}
