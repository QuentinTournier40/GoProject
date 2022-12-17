package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"goproject/cmd/pub/config"
	"log"
	"time"
)

func main() {

	configuration := config.GetConfig("Pressure")
	address := configuration.ADDRESS
	port := configuration.PORT
	qos := configuration.QOS
	clientId := configuration.CLIENT_ID
	delay := configuration.DELAY

	fmt.Println(delay)

	topic := "capteurs"
	client := connect(address+":"+port, clientId)

	for {
		now := time.Now()
		msg := "2 LYN PRESSURE " + "valeur" + " " + now.Format("2006-02-01-15-04-05")
		client.Publish(topic, qos, false, msg)
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
