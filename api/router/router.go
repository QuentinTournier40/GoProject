package router

import (
	"github.com/gorilla/mux"
	"goproject/api/captorService"
	"log"
	"net/http"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/get/data-by-iata-code/{iataCode}", captorService.GetDataByIataCode).Methods("GET")
	myRouter.HandleFunc("/get/data-by-iata-code-and-number/{iataCode}/{number}", captorService.GetDataByIataCodeForXData).Methods("GET")
	myRouter.HandleFunc("/get/data-by-captorName/{captorName}", captorService.GetDataByCaptor).Methods("GET")
	myRouter.HandleFunc("/get/data-between-dates/{captorName}/{start}/{end}", captorService.GetDataBetweenDates).Methods("GET")
	myRouter.HandleFunc("/get/average-data/{date}", captorService.GetAverageByDate).Methods("GET")
	myRouter.HandleFunc("/get/data/{iataCode}/{captorName}", captorService.GetDataByIataCodeAndCaptor).Methods("GET")
	myRouter.HandleFunc("/get/data/{iataCode}/{captorName}/{year}", captorService.GetDataByIataCodeAndCaptorAndYear).Methods("GET")
	myRouter.HandleFunc("/get/data/{iataCode}/{captorName}/{year}/{month}", captorService.GetDataByIataCodeAndCaptorAndYearAndMonth).Methods("GET")
	myRouter.HandleFunc("/get/data/{iataCode}/{captorName}/{year}/{month}/{day}", captorService.GetDataByIataCodeAndCaptorAndYearAndMonthAndDay).Methods("GET")
	log.Println("Server listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
