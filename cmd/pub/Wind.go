package main

import (
	"fmt"
)

func main() {
	topic := "test"
	msg := "Hello world!"
	client := connect("tcp://localhost:1883", "publish")
	client.Publish(topic, 0, false, msg)
	fmt.Println("==============================\n" +
		"Message envoyé au sujet: " + topic +
		"\n==============================\n")
}

/*
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

*/
