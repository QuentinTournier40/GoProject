package main

import "goproject/api/router"

// @Version 1.0.0
// @Title Temperature/Pressure/Wind of airports API
// @Description l'API expose les données generés par nos publishers
// @ContactName TOURNIER Quentin
// @ContactEmail qttournier@gmail.com
func main() {
	router.HandleRequests()
}
