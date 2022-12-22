package Captors

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"goproject/cmd/pub/config"
	"log"
	"math/rand"
	"time"
)

func RunCaptor(captorFileName, captorFullNameUpperCase, captorId string, minValue, maxValue float64) {
	// GENERATE RANDOM SEED
	rand.Seed(time.Now().UnixNano())
	// CONFIG
	configuration := config.GetConfig(captorFileName)
	address := configuration.ADDRESS
	port := configuration.PORT
	qos := configuration.QOS
	clientId := configuration.CLIENT_ID
	delay := configuration.DELAY

	topic := "capteurs"
	client := connectClient(address+":"+port, clientId, delay)

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
			msg := captorId + " " + value + " " + captorFullNameUpperCase + " " + fmt.Sprintf("%.1f", tabValue[key]) + " " + now.Format("2006-02-01-15-04-05")
			client.Publish(topic, qos, false, msg)
			fmt.Println(msg)
		}
		time.Sleep(time.Duration(delay) * time.Second)
	}
}

func createClientOptions(brokerURI, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientId)
	return opts
}

func connectClient(brokerURI, clientId string, delay int) mqtt.Client {
	fmt.Println("Trying to connect (" + brokerURI + ", " + clientId + ")...")
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(time.Duration(delay) * time.Second) {
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
