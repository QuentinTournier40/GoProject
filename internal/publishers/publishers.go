package publishers

import (
	"fmt"
	config2 "goproject/internal/config"
	"goproject/internal/pubSubMethods"
	"math/rand"
	"strconv"
	"time"
)

func RunPublisher(captorFullNameUpperCase string, captorId int, minValue, maxValue float64) {
	// GENERATE RANDOM SEED
	rand.Seed(time.Now().UnixNano())

	configuration := config2.GetConfig()
	clientId := ""
	switch captorId {
	case 1:
		clientId = configuration.PRESSURE.CLIENT_ID
	case 2:
		clientId = configuration.TEMPERATURE.CLIENT_ID
	case 3:
		clientId = configuration.WIND.CLIENT_ID
	}

	topic := "capteurs"
	client := pubSubMethods.Connect(configuration.ADDRESS+":"+configuration.PORT, clientId, configuration.DELAY)

	mapIata := config2.CODE_IATA

	var tabValue []float64

	// SET FIRST RANDOM VALUE
	for range mapIata {
		tabValue = append(tabValue, minValue+rand.Float64()*(maxValue-minValue))
	}

	for {
		for key, value := range mapIata {
			tabValue[key] = generateCoherenteValue(tabValue[key], minValue, maxValue)
			now := time.Now()
			msg := strconv.FormatInt(int64(3*key+captorId), 10) + " " + value + " " + captorFullNameUpperCase + " " + fmt.Sprintf("%.1f", tabValue[key]) + " " + now.Format("2006-01-02-15-04-05")
			client.Publish(topic, configuration.QOS, false, msg)
		}
		time.Sleep(time.Duration(configuration.DELAY) * time.Second)
	}
}

func generateCoherenteValue(value float64, min, max float64) float64 {
	offset := rand.Float64()*6 - 3
	result := value + offset

	if result < min {
		result = min
	} else if result > max {
		result = max
	}
	return result
}
