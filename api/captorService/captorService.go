package captorService

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"goproject/bdd"
	"goproject/config"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ------------------------------ TYPES JSON ------------------------------

type AllCaptors struct {
	IATA        string     `json:"iata"`
	PRESSURE    []*Measure `json:"pressure"`
	TEMPERATURE []*Measure `json:"temperature"`
	WIND        []*Measure `json:"wind"`
}

type Measure struct {
	DATE  string `json:"date"`
	VALUE string `json:"value"`
}

type Captor struct {
	CAPTORNAME string  `json:"captorName"`
	VALUES     []*Iata `json:"values"`
}

type Iata struct {
	IATA     string     `json:"iata"`
	MEASURES []*Measure `json:"measures"`
}

type CaptorAndIata struct {
	IATA       string     `json:"iata"`
	CAPTORNAME string     `json:"captorName"`
	VALUES     []*Measure `json:"values"`
}

type BetweenDate struct {
	START      string  `json:"start"`
	END        string  `json:"end"`
	CAPTORNAME string  `json:"captorName"`
	VALUES     []*Iata `json:"values"`
}

type DateAndAllCaptors struct {
	DATE        string  `json:"date"`
	PRESSURE    float64 `json:"pressure"`
	TEMPERATURE float64 `json:"temperature"`
	WIND        float64 `json:"wind"`
}

// ------------------------------ SERVICES ------------------------------

func GetDataByIataCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	iataCode := strings.ToUpper(vars["iataCode"])

	var pressureMeasures []*Measure
	var temperatureMeasures []*Measure
	var windMeasures []*Measure

	keysPressure := bdd.GetAllKeyRegex(iataCode + "/PRESSURE/*")
	keysTemperature := bdd.GetAllKeyRegex(iataCode + "/TEMPERATURE/*")
	keysWind := bdd.GetAllKeyRegex(iataCode + "/WIND/*")

	for _, key := range keysPressure {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		pressureMeasures = append(pressureMeasures, &Measure{DATE: splitKey[2], VALUE: data})
	}
	for _, key := range keysTemperature {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		temperatureMeasures = append(temperatureMeasures, &Measure{DATE: splitKey[2], VALUE: data})
	}
	for _, key := range keysWind {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		windMeasures = append(windMeasures, &Measure{DATE: splitKey[2], VALUE: data})
	}
	p, _ := json.Marshal(AllCaptors{IATA: iataCode, PRESSURE: pressureMeasures, TEMPERATURE: temperatureMeasures, WIND: windMeasures})
	w.Write(p)
}

func GetDataByCaptor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	captorName := strings.ToUpper(vars["captorName"])

	var iataData []*Iata
	var measures []*Measure

	mapIata := config.CODE_IATA

	for _, iata := range mapIata {
		measures = nil
		keys := bdd.GetAllKeyRegex(strings.ToUpper(iata) + "/" + captorName + "/*")
		for _, key := range keys {
			data := bdd.GetValue(key)
			splitKey := strings.Split(key, "/")
			measures = append(measures, &Measure{DATE: splitKey[2], VALUE: data})
		}
		iataData = append(iataData, &Iata{IATA: iata, MEASURES: measures})
	}

	p, _ := json.Marshal(Captor{CAPTORNAME: captorName, VALUES: iataData})
	w.Write(p)
}

func GetDataByIataCodeAndCaptor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	iataCode := strings.ToUpper(vars["iataCode"])
	captorName := strings.ToUpper(vars["captorName"])

	var measures []*Measure

	keys := bdd.GetAllKeyRegex(iataCode + "/" + captorName + "/*")

	for _, key := range keys {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		measures = append(measures, &Measure{DATE: splitKey[2], VALUE: data})
	}
	p, _ := json.Marshal(CaptorAndIata{IATA: iataCode, CAPTORNAME: captorName, VALUES: measures})
	w.Write(p)
}

func GetDataByIataCodeAndCaptorAndYear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	iataCode := strings.ToUpper(vars["iataCode"])
	captorName := strings.ToUpper(vars["captorName"])
	year := vars["year"]

	var measures []*Measure

	keys := bdd.GetAllKeyRegex(iataCode + "/" + captorName + "/" + year + "-*")

	for _, key := range keys {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		measures = append(measures, &Measure{DATE: splitKey[2], VALUE: data})
	}
	p, _ := json.Marshal(CaptorAndIata{IATA: iataCode, CAPTORNAME: captorName, VALUES: measures})
	w.Write(p)
}

func GetDataByIataCodeAndCaptorAndYearAndMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	iataCode := strings.ToUpper(vars["iataCode"])
	captorName := strings.ToUpper(vars["captorName"])
	year := vars["year"]
	month := vars["month"]

	var measures []*Measure

	keys := bdd.GetAllKeyRegex(iataCode + "/" + captorName + "/" + year + "-" + month + "-*")

	for _, key := range keys {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		measures = append(measures, &Measure{DATE: splitKey[2], VALUE: data})
	}
	p, _ := json.Marshal(CaptorAndIata{IATA: iataCode, CAPTORNAME: captorName, VALUES: measures})
	w.Write(p)
}

func GetDataByIataCodeAndCaptorAndYearAndMonthAndDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	iataCode := strings.ToUpper(vars["iataCode"])
	captorName := strings.ToUpper(vars["captorName"])
	year := vars["year"]
	month := vars["month"]
	day := vars["day"]

	var measures []*Measure

	keys := bdd.GetAllKeyRegex(iataCode + "/" + captorName + "/" + year + "-" + month + "-" + day + "-*")

	for _, key := range keys {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		measures = append(measures, &Measure{DATE: splitKey[2], VALUE: data})
	}
	p, _ := json.Marshal(CaptorAndIata{IATA: iataCode, CAPTORNAME: captorName, VALUES: measures})
	w.Write(p)
}

func GetDataBetweenDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	captorName := strings.ToUpper(vars["captorName"])
	start := vars["start"]
	end := vars["end"]

	var iataData []*Iata
	var measures []*Measure
	mapIata := config.CODE_IATA

	startDate, _ := time.Parse("2006-01-02-15", start)
	endDate, _ := time.Parse("2006-01-02-15", end)

	var keys []string
	for date := startDate; date != endDate; date = date.Add(time.Hour) {
		dateStr := date.Format("2006-01-02-15")
		keys = append(keys, bdd.GetAllKeyRegex("*/"+captorName+"/"+dateStr+"*")...)
	}

	for _, iata := range mapIata {
		measures = nil
		for _, key := range keys {
			splitKey := strings.Split(key, "/")
			if splitKey[0] == iata {
				data := bdd.GetValue(key)
				measures = append(measures, &Measure{DATE: splitKey[2], VALUE: data})
			}
			fmt.Println(len(keys))
		}
		iataData = append(iataData, &Iata{IATA: iata, MEASURES: measures})
	}

	p, _ := json.Marshal(BetweenDate{START: start, END: end, CAPTORNAME: captorName, VALUES: iataData})
	w.Write(p)
}

func GetAverageByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	date := vars["date"]

	pressureSum := 0.0
	temperatureSum := 0.0
	windSum := 0.0

	dateDate, _ := time.Parse("2006-01-02", date)
	dateFormat := dateDate.Format("2006-01-02")

	keysPressure := bdd.GetAllKeyRegex("*/PRESSURE/" + dateFormat + "*")
	keysTemperature := bdd.GetAllKeyRegex("*/TEMPERATURE/" + dateFormat + "*")
	keysWind := bdd.GetAllKeyRegex("*/WIND/" + dateFormat + "*")

	for _, key := range keysPressure {
		data := bdd.GetValue(key)
		pressure, _ := strconv.ParseFloat(data, 64)
		pressureSum += pressure
	}
	for _, key := range keysTemperature {
		data := bdd.GetValue(key)
		temperature, _ := strconv.ParseFloat(data, 64)
		temperatureSum += temperature
	}
	for _, key := range keysWind {
		data := bdd.GetValue(key)
		wind, _ := strconv.ParseFloat(data, 64)
		windSum += wind
	}

	averagePressure := pressureSum / float64(len(keysPressure))
	averageTemperature := temperatureSum / float64(len(keysTemperature))
	averageWind := windSum / float64(len(keysWind))

	if math.IsNaN(averagePressure) {
		averagePressure = math.SmallestNonzeroFloat64
	}
	if math.IsNaN(averageTemperature) {
		averageTemperature = math.SmallestNonzeroFloat64
	}
	if math.IsNaN(averageWind) {
		averageWind = math.SmallestNonzeroFloat64
	}

	j, _ := json.Marshal(DateAndAllCaptors{DATE: date, WIND: averageWind, TEMPERATURE: averageTemperature, PRESSURE: averagePressure})
	w.Write(j)
}
