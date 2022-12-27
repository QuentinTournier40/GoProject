package Subscribers

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"goproject/bdd"
	"goproject/cmd/PubSubMethods"
	"goproject/config"
	"strings"
	"sync"
)

func RunSubscriber(clientId string, isForApi bool) {
	configuration := config.GetConfig()
	topic := "capteurs"
	client := PubSubMethods.Connect(configuration.ADDRESS+":"+configuration.PORT, clientId, configuration.DELAY)

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
			fmt.Println("J'ecris dans le subscriber_csv")
		}
	})
	wg.Wait()
}
