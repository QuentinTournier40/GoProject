package Subscribers

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"goproject/bdd"
	"goproject/cmd/PubSubMethods"
	"strings"
	"sync"
)

func RunSubscriber(clientId string, isForApi bool) {
	topic := "capteurs"
	// TODO REMOVE 3 AND ADD REAL DELAY FROM CONFIG
	client := PubSubMethods.Connect("tcp://localhost:1883", clientId, 3)

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
