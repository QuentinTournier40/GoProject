package router

import (
	"github.com/gorilla/mux"
	"goproject/api/captorService"
	"log"
	"net/http"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/get/all-data", captorService.GetDataFromAllCaptors).Methods("GET")
	myRouter.HandleFunc("/get/data-by-iata-code/{iataCode}", captorService.GetDataByIataCode).Methods("GET")
	myRouter.HandleFunc("/get/data-by-captorService/{captorName}", captorService.GetDataByCaptor).Methods("GET")
	myRouter.HandleFunc("/get/{iataCode}/{captorName}", captorService.GetDataByIataCodeAndCaptor).Methods("GET")
	myRouter.HandleFunc("/get/{iataCode}/{captorName}/{year}", captorService.GetDataByIataCodeAndCaptorAndYear).Methods("GET")
	myRouter.HandleFunc("/get/{iataCode}/{captorName}/{year}/{month}", captorService.GetDataByIataCodeAndCaptorAndYearAndMonth).Methods("GET")
	myRouter.HandleFunc("/get/{iataCode}/{captorName}/{year}/{month}/{day}", captorService.GetDataByIataCodeAndCaptorAndYearAndMonthAndDay).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
