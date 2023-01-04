package Subscribers

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-co-op/gocron"
	"golang.org/x/exp/maps"
	"goproject/internal/PubSubMethods"
	"goproject/internal/bdd"
	"goproject/internal/config"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func RunSubscriber(clientId string, isForApi bool) {
	configuration := config.GetConfig()
	topic := "capteurs"
	client := PubSubMethods.Connect(configuration.ADDRESS+":"+configuration.PORT, clientId, configuration.DELAY)

	var wg sync.WaitGroup
	wg.Add(1)

	var mapCsv = map[string][]string{}

	if isForApi {
		client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			value := strings.Split(string(msg.Payload()), " ")
			key := value[1] + "/" + value[2] + "/" + value[4]
			bdd.SetValue(key, value[3])
		})
	} else {
		job := gocron.NewScheduler(time.UTC)
		// TODO VERIFIER QUE L'HEURE DE BASE C'EST BIEN MINUIT
		job.Every(1).Day().Do(func() {
			createCsvFiles(mapCsv)
		})
		job.StartAsync()

		client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			value := strings.Split(string(msg.Payload()), " ")
			tab := mapCsv[value[1]+"_"+value[2]+"_"+value[4][:10]]
			tab = append(tab, strings.Join(value, ";"))
			mapCsv[value[1]+"_"+value[2]+"_"+value[4][:10]] = tab
		})
	}
	wg.Wait()
}

func createCsvFiles(mapCsv map[string][]string) {
	for key, value := range mapCsv {
		csvFile, err := os.Create("./csv/" + key + ".csv")
		if err != nil {
			log.Fatalln("Failed creating file : %s", err)
		}
		csvFile.WriteString("Identifiant capteur;Code IATA;Intitule du capteur;Valeur;Date;\n")
		for _, data := range value {
			csvFile.WriteString(data + ";\n")
		}
		csvFile.Close()
	}
	maps.Clear(mapCsv)
}
