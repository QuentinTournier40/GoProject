package captorService

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"goproject/bdd"
	"net/http"
	"strings"
)

type Datas struct {
	DATA []*Data `json:"values"`
}

// TODO VOIR OU FOUTRE TOUS CES STRUCT

type Data struct {
	IATA         string `json:"iata"`
	MEASURETYPE  string `json:"measuretype"`
	MEASUREVALUE string `json:"measurevalue"`
	MEASUREDATE  string `json:"measuredate"`
}

type DatasIataCode struct {
	IATA        string         `json:"iata"`
	PRESSURE    []*DataMeasure `json:"pressure"`
	TEMPERATURE []*DataMeasure `json:"temperature"`
	WIND        []*DataMeasure `json:"wind"`
}

type DataMeasure struct {
	DATE  string `json:"date"`
	VALUE string `json:"value"`
}

func GetDataFromAllCaptors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dataValues []*Data
	keys := bdd.GetAllKey()
	for _, key := range keys {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		dataValues = append(dataValues, &Data{IATA: splitKey[0], MEASURETYPE: splitKey[1], MEASUREVALUE: data, MEASUREDATE: splitKey[2]})
	}
	p, _ := json.Marshal(Datas{DATA: dataValues})
	w.Write(p)
}

func GetDataByIataCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	iataCode := vars["iataCode"]
	var pressureData []*DataMeasure
	var temperatureData []*DataMeasure
	var windData []*DataMeasure
	keysPressure := bdd.GetAllKeyRegex(iataCode + "/PRESSURE")
	keysTemperature := bdd.GetAllKeyRegex(iataCode + "/TEMPERATURE")
	keysWind := bdd.GetAllKeyRegex(iataCode + "/WIND")
	for _, key := range keysPressure {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		pressureData = append(pressureData, &DataMeasure{DATE: splitKey[2], VALUE: data})
	}
	for _, key := range keysTemperature {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		temperatureData = append(temperatureData, &DataMeasure{DATE: splitKey[2], VALUE: data})
	}
	for _, key := range keysWind {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		windData = append(windData, &DataMeasure{DATE: splitKey[2], VALUE: data})
	}
	p, _ := json.Marshal(DatasIataCode{IATA: iataCode, PRESSURE: pressureData, TEMPERATURE: temperatureData, WIND: windData})
	w.Write(p)
}

func GetDataByCaptor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get data from one captorService")
}

func GetDataByIataCodeAndCaptor(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get data from one captorService but only in one airport")
}

func GetDataByIataCodeAndCaptorAndYear(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get data from one captorService but only in one airport on specific year")
}

func GetDataByIataCodeAndCaptorAndYearAndMonth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get data from one captorService but only in one airport on specific year and month")
}

func GetDataByIataCodeAndCaptorAndYearAndMonthAndDay(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get data from one captorService but only in one airport on specific year, and day")
}
