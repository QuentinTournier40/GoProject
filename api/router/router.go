package router

import (
	"github.com/gorilla/mux"
	"goproject/api/apiService"
	"log"
	"net/http"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/iata/{code}", apiService.GetDataByIataCode).Methods("GET")
	myRouter.HandleFunc("/iata/{code}/number/{number}", apiService.GetDataByIataCodeForXData).Methods("GET")
	myRouter.HandleFunc("/sensor/{sensorName}", apiService.GetDataByCaptor).Methods("GET")
	myRouter.HandleFunc("/iata/{code}/sensor/{sensorName}", apiService.GetDataByIataCodeAndCaptor).Methods("GET")
	myRouter.HandleFunc("/sensor/{sensorName}/between/{start_date}/to/{end_date}", apiService.GetDataBetweenDates).Methods("GET")
	myRouter.HandleFunc("/averages/{date}", apiService.GetAverageByDate).Methods("GET")
	log.Println("Server listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
