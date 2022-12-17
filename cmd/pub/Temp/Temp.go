package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"goproject/cmd/pub/config"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	configuration := config.GetConfig("Temp")
	address := configuration.ADDRESS
	port := configuration.PORT
	qos := configuration.QOS
	clientId := configuration.CLIENT_ID
	//delay := configuration.DELAY

	topic := "capteurs"
	client := connect(address+":"+port, clientId)

	mapIata := config.CODE_IATA

	var tableauTemperature []float64

	for range mapIata {
		tableauTemperature = append(tableauTemperature, -15+rand.Float64()*(30 - -15))
	}

	for {
		for key, value := range mapIata {
			tableauTemperature[key] = generateCoherenteValue(tableauTemperature[key])
			now := time.Now()
			msg := "1 " + value + " TEMPERATURE " + fmt.Sprintf("%.1f", tableauTemperature[key]) + " " + now.Format("2006-02-01-15-04-05")
			client.Publish(topic, qos, false, msg)
		}

		time.Sleep(3 * time.Second)

	}
}

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientId)
	return opts
}

func connect(brokerURI string, clientId string) mqtt.Client {
	fmt.Println("Trying to connect (" + brokerURI + ", " + clientId + ")...")
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
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
