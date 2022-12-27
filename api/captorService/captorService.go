package captorService

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"goproject/bdd"
	"net/http"
	"strconv"
	"strings"
	"time"
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

type DataCaptor struct {
	CAPTORNAME string                 `json:"captorName"`
	VALUES     []*DataMeasureCodeIata `json:"values"`
}

type DataMeasureCodeIata struct {
	IATA  string `json:"iata"`
	DATE  string `json:"date"`
	VALUE string `json:"value"`
}

type DataCaptorIataCode struct {
	IATA       string         `json:"iata"`
	CAPTORNAME string         `json:"captorName"`
	VALUES     []*DataMeasure `json:"values"`
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
	keysPressure := bdd.GetAllKeyRegex(iataCode + "/PRESSURE/*")
	keysTemperature := bdd.GetAllKeyRegex(iataCode + "/TEMPERATURE/*")
	keysWind := bdd.GetAllKeyRegex(iataCode + "/WIND/*")
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
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	captorName := vars["captorName"]
	var captorData []*DataMeasureCodeIata
	var keys []string
	switch strings.ToLower(captorName) {
	case "pressure":
		keys = bdd.GetAllKeyRegex("*/PRESSURE/*")
	case "temperature":
		keys = bdd.GetAllKeyRegex("*/TEMPERATURE/*")
	case "wind":
		keys = bdd.GetAllKeyRegex("*/WIND/*")
	}

	for _, key := range keys {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		captorData = append(captorData, &DataMeasureCodeIata{IATA: splitKey[0], DATE: splitKey[2], VALUE: data})
	}
	p, _ := json.Marshal(DataCaptor{CAPTORNAME: captorName, VALUES: captorData})
	w.Write(p)
}

func GetDataByIataCodeAndCaptor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	iataCode := vars["iataCode"]
	captorName := vars["captorName"]
	var captorData []*DataMeasure
	var keys []string
	switch strings.ToLower(captorName) {
	case "pressure":
		keys = bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/PRESSURE/*")
	case "temperature":
		keys = bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/TEMPERATURE/*")
	case "wind":
		keys = bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/WIND/*")
	}

	for _, key := range keys {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		captorData = append(captorData, &DataMeasure{DATE: splitKey[2], VALUE: data})
	}
	p, _ := json.Marshal(DataCaptorIataCode{IATA: iataCode, CAPTORNAME: captorName, VALUES: captorData})
	w.Write(p)

}

func GetDataByIataCodeAndCaptorAndYear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	iataCode := vars["iataCode"]
	captorName := vars["captorName"]
	year := vars["year"]
	var captorData []*DataMeasure
	var keys []string
	switch strings.ToLower(captorName) {
	case "pressure":
		keys = bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/PRESSURE/" + year + "-*")
	case "temperature":
		keys = bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/TEMPERATURE/" + year + "-*")
	case "wind":
		keys = bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/WIND/" + year + "-*")
	}

	for _, key := range keys {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		captorData = append(captorData, &DataMeasure{DATE: splitKey[2], VALUE: data})
	}
	p, _ := json.Marshal(DataCaptorIataCode{IATA: iataCode, CAPTORNAME: captorName, VALUES: captorData})
	w.Write(p)
}

func GetDataByIataCodeAndCaptorAndYearAndMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	iataCode := vars["iataCode"]
	captorName := vars["captorName"]
	year := vars["year"]
	month := vars["month"]
	var captorData []*DataMeasure
	var keys []string
	switch strings.ToLower(captorName) {
	case "pressure":
		keys = bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/PRESSURE/" + year + "-" + month + "-*")
	case "temperature":
		keys = bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/TEMPERATURE/" + year + "-" + month + "-*")
	case "wind":
		keys = bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/WIND/" + year + "-" + month + "-*")
	}

	for _, key := range keys {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		captorData = append(captorData, &DataMeasure{DATE: splitKey[2], VALUE: data})
	}
	p, _ := json.Marshal(DataCaptorIataCode{IATA: iataCode, CAPTORNAME: captorName, VALUES: captorData})
	w.Write(p)
}

func GetDataByIataCodeAndCaptorAndYearAndMonthAndDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	iataCode := vars["iataCode"]
	captorName := vars["captorName"]
	year := vars["year"]
	month := vars["month"]
	day := vars["day"]
	var captorData []*DataMeasure
	var keys []string
	switch strings.ToLower(captorName) {
	case "pressure":
		keys = bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/PRESSURE/" + year + "-" + month + "-" + day + "-*")
	case "temperature":
		keys = bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/TEMPERATURE/" + year + "-" + month + "-" + day + "-*")
	case "wind":
		keys = bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/WIND/" + year + "-" + month + "-" + day + "-*")
	}

	for _, key := range keys {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		captorData = append(captorData, &DataMeasure{DATE: splitKey[2], VALUE: data})
	}
	p, _ := json.Marshal(DataCaptorIataCode{IATA: iataCode, CAPTORNAME: captorName, VALUES: captorData})
	w.Write(p)
}

func GetDataBetweenDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	captorName := vars["captorName"]
	start := vars["start"]
	end := vars["end"]
	//var captorData []*DataMeasure

	startDate, _ := time.Parse("2006-01-02-15", start)
	endDate, _ := time.Parse("2006-01-02-15", end)

	var keys []string
	for d := startDate; d.After(endDate) == false; d = d.Add(time.Hour) {
		t := d.Format("2006-01-02-15")
		keys = append(keys, bdd.GetAllKeyRegex("*/"+strings.ToUpper(captorName)+"/"+t+"*")...)
	}

	for _, key := range keys {
		fmt.Fprintf(w, key)
	}
}

func GetAverageByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	date := vars["date"]
	var pressureData float64 = 0
	var temperatureData float64 = 0
	var windData float64 = 0

	fmt.Println(date)
	dateDate, _ := time.Parse("2006-01-02", date)

	fmt.Println(dateDate.String())

	fmt.Println("*/PRESSURE/" + dateDate.Format("2006-01-02") + "*")
	keysPressure := bdd.GetAllKeyRegex("*/PRESSURE/" + dateDate.Format("2006-01-02") + "*")
	keysTemperature := bdd.GetAllKeyRegex("*/TEMPERATURE/" + dateDate.Format("2006-01-02") + "*")
	keysWind := bdd.GetAllKeyRegex("*/WIND/" + dateDate.Format("2006-01-02") + "*")
	for _, key := range keysPressure {
		data := bdd.GetValue(key)
		chif, _ := strconv.ParseFloat(data, 64)
		pressureData += chif
		//pressureData = append(pressureData, &DataMeasure{DATE: splitKey[2], VALUE: data})
	}
	for _, key := range keysTemperature {
		data := bdd.GetValue(key)
		chif, _ := strconv.ParseFloat(data, 64)
		temperatureData += chif
		//temperatureData = append(temperatureData, &DataMeasure{DATE: splitKey[2], VALUE: data})
	}
	for _, key := range keysWind {
		data := bdd.GetValue(key)
		chif, _ := strconv.ParseFloat(data, 64)
		windData += chif

		//windData = append(windData, &DataMeasure{DATE: splitKey[2], VALUE: data})
	}
	fmt.Fprintf(w, fmt.Sprintf("%f", pressureData/float64(len(keysPressure)))+"\n")
	fmt.Fprintf(w, fmt.Sprintf("%f", temperatureData/float64(len(keysTemperature)))+"\n")
	fmt.Fprintf(w, fmt.Sprintf("%f", windData/float64(len(keysWind)))+"\n")
}
