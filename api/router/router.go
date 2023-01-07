package router

import (
	"github.com/gorilla/mux"
	"goproject/api/apiService"
	"log"
	"net/http"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/iata/{code}", apiService.GetDataByIataCode).Methods("GET")                                              //  /get/data-by-iata-code/{iataCode}
	myRouter.HandleFunc("/iata/{code}/number/{number}", apiService.GetDataByIataCodeForXData).Methods("GET")                      // /iata/{code}/number/{num} /get/data-by-iata-code-and-number/{iataCode}/{number}
	myRouter.HandleFunc("/sensor/{sensorName}", apiService.GetDataByCaptor).Methods("GET")                                        // /sensor/{sensorName} /get/data-by-captorName/{captorName}
	myRouter.HandleFunc("/iata/{code}/sensor/{sensorName}", apiService.GetDataByIataCodeAndCaptor).Methods("GET")                 //  /get/data/{iataCode}/{captorName}
	myRouter.HandleFunc("/sensor/{sensorName}/between/{start_date}/to/{end_date}", apiService.GetDataBetweenDates).Methods("GET") // /sensor/{sensorName}/between/{start_date}/to/{end_date} /get/data-between-dates/{captorName}/{start}/{end}
	myRouter.HandleFunc("/averages/{date}", apiService.GetAverageByDate).Methods("GET")                                           //  /get/average-data/{date}
	log.Println("Server listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
