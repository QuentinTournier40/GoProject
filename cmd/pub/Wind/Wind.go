package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"goproject/cmd/pub/config"
	"log"
	"math"
	"math/rand"
	"time"
)

func main() {

	configuration := config.GetConfig("Wind")
	address := configuration.ADDRESS
	port := configuration.PORT
	qos := configuration.QOS
	clientId := configuration.CLIENT_ID
	//delay := configuration.DELAY

	topic := "capteurs"
	client := connect(address+":"+port, clientId)

	for {
		now := time.Now()
		msg := "3 LYN WIND " + "valeur" + " " + now.Format("2006-02-01-15-04-05")
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

// GenereTemperature simule la lecture de température d'un capteur de température
// en générant une valeur aléatoire autour de la valeur de référence, avec une incertitude donnée
func GenereTemperature(min, max float64, incertitude float64, t float64) float64 {
	// Initialise le générateur de nombres aléatoires avec un seed tiré de l'horloge
	// de l'ordinateur. Cela permet de générer des nombres aléatoires différents à chaque exécution
	// du programme.
	rand.Seed(time.Now().UnixNano())

	// Génère une valeur aléatoire dans l'intervalle [min, max]
	reference := rand.Float64()*(max-min) + min

	// Génère un nombre aléatoire dans l'intervalle [-incertitude, incertitude]
	alea := rand.Float64()*2*incertitude - incertitude

	// Ajoute l'incertitude aléatoire à la valeur de référence pour obtenir la valeur de température simulée
	temperature := reference + alea

	// Modifie la température en fonction du temps en utilisant une fonction sinusoïdale
	temperature += 5 * math.Sin(t/10)

	return temperature
}
