package main

import "goproject/api/router"

// @Version 1.0.0
// @Title API de données de capteurs d'aeroports
// @Description l'API expose les données générées par nos publishers
func main() {
	router.HandleRequests()
}
