package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"goproject/bdd"
	"log"
	"strings"
	"sync"
	"time"
)

func main() {
	topic := "capteurs"
	client := connect("tcp://localhost:1883", "test-client")

	var wg sync.WaitGroup
	wg.Add(1)

	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		//Get the Kuzzle response
		//fmt.Println(string(msg.Payload()))
		table := strings.Split(string(msg.Payload()), " ")

		fmt.Println(table[1])
		key := table[1] + "/" + table[2] + "/" + table[4]
		bdd.SetValue(key, table[3])

	})

	wg.Wait()
}

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	// AddBroker adds a broker URI to the list of brokers to be used.
	// The format should be "scheme://host:port"
	opts.AddBroker(brokerURI)
	//
	//opts.SetUsername(user)
	////
	//opts.SetPassword(password)
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
