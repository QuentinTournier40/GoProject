package Subscribers

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"goproject/bdd"
	"log"
	"strings"
	"sync"
	"time"
)

func RunSubscriber(clientId string, isForApi bool) {
	topic := "capteurs"
	client := connect("tcp://localhost:1883", clientId)

	var wg sync.WaitGroup
	wg.Add(1)

	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		if isForApi {
			table := strings.Split(string(msg.Payload()), " ")
			fmt.Println(table[1])
			key := table[1] + "/" + table[2] + "/" + table[4]
			bdd.SetValue(key, table[3])
		} else {
			// ECRITURE DANS LE FICHIER CSV
			fmt.Println("J'ecris dans le csv")
		}
	})
	wg.Wait()
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
