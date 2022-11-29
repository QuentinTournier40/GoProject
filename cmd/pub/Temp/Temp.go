package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"goproject/cmd/pub/config"
	"log"
	"math/rand"
	"time"
)

func main() {
	type DataJson struct {
		id     int
		iata   string
		nature string
		value  float64
		date   string
	}

	configuration := config.GetConfig("Temp")
	address := configuration.ADDRESS
	port := configuration.PORT
	qos := configuration.QOS
	clientId := configuration.CLIENT_ID
	//delay := configuration.DELAY

	topic := "test"
	client := connect(address+":"+port, clientId)

	tempGenerate := -15 + rand.Float64()*(30 - -15)

	data := DataJson{1, "LYN", "Temperature", 30.2, "YYYY-MM-DD"}
	json, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	for {
		msg := "Temperature actuelle: "
		tempGenerate := generateCoherenteValue(tempGenerate)
		msg += fmt.Sprintf("%v", tempGenerate)
		client.Publish(topic, qos, false, json)
		fmt.Println("==============================\n" +
			"Message envoyé au sujet: " + topic +
			"\n==============================\n")
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
	println(val)
	return val
}
