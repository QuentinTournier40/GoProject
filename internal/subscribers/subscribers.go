package subscribers

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-co-op/gocron"
	"golang.org/x/exp/maps"
	"goproject/internal/bdd"
	"goproject/internal/config"
	"goproject/internal/pubSubMethods"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func RunSubscriber(clientId string, isForApi bool) {
	configuration := config.GetConfig()
	topic := "capteurs"
	client := pubSubMethods.Connect(configuration.ADDRESS+":"+configuration.PORT, clientId, configuration.DELAY)

	var wg sync.WaitGroup
	wg.Add(1)

	if isForApi {
		subscribeApi(client, topic)
	} else {
		subscribeCsv(client, topic)
	}
	wg.Wait()
}

func subscribeApi(client mqtt.Client, topic string) {
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		strMsg := strings.Split(string(msg.Payload()), " ")
		date, _ := time.Parse("2006-01-02-15-04-05", strMsg[4])
		dataRedis := strMsg[4] + ":" + strMsg[3] + ":" + strMsg[0]
		bdd.AddToSortedSet(strMsg[1]+"/"+strMsg[2], date.Unix(), dataRedis)
	})
}

func subscribeCsv(client mqtt.Client, topic string) {
	var mapCsv = map[string][]string{}

	job := gocron.NewScheduler(time.UTC)
	job.Every(1).Day().Do(func() {
		createCsvFiles(mapCsv)
	})
	job.StartAsync()

	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		strMsg := strings.Split(string(msg.Payload()), " ")
		tab := mapCsv[strMsg[1]+"_"+strMsg[2]+"_"+strMsg[4][:10]]
		tab = append(tab, strings.Join(strMsg, ";"))
		mapCsv[strMsg[1]+"_"+strMsg[2]+"_"+strMsg[4][:10]] = tab
	})
}

func createCsvFiles(mapCsv map[string][]string) {
	for key, value := range mapCsv {
		csvFile, err := os.Create("../csv/" + key + ".csv")
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
