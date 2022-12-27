package captorService

import (
	"encoding/json"
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
	iataCode := vars["iataCode"]

	var pressureData []*Measure
	var temperatureData []*Measure
	var windData []*Measure

	keysPressure := bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/PRESSURE/*")
	keysTemperature := bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/TEMPERATURE/*")
	keysWind := bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/WIND/*")

	for _, key := range keysPressure {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		pressureData = append(pressureData, &Measure{DATE: splitKey[2], VALUE: data})
	}
	for _, key := range keysTemperature {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		temperatureData = append(temperatureData, &Measure{DATE: splitKey[2], VALUE: data})
	}
	for _, key := range keysWind {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		windData = append(windData, &Measure{DATE: splitKey[2], VALUE: data})
	}
	p, _ := json.Marshal(AllCaptors{IATA: iataCode, PRESSURE: pressureData, TEMPERATURE: temperatureData, WIND: windData})
	w.Write(p)
}

func GetDataByCaptor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	captorName := vars["captorName"]

	var iataData []*Iata
	var mesure []*Measure
	var keys []string

	mapIata := config.CODE_IATA

	for _, iata := range mapIata {
		mesure = nil
		keys = bdd.GetAllKeyRegex(strings.ToUpper(iata) + "/" + strings.ToUpper(captorName) + "/*")
		for _, key := range keys {
			data := bdd.GetValue(key)
			splitKey := strings.Split(key, "/")
			mesure = append(mesure, &Measure{DATE: splitKey[2], VALUE: data})
		}
		iataData = append(iataData, &Iata{IATA: iata, MEASURES: mesure})
	}

	p, _ := json.Marshal(Captor{CAPTORNAME: captorName, VALUES: iataData})
	w.Write(p)
}

func GetDataByIataCodeAndCaptor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	iataCode := vars["iataCode"]
	captorName := vars["captorName"]

	var captorData []*Measure

	keys := bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/" + strings.ToUpper(captorName) + "/*")

	for _, key := range keys {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		captorData = append(captorData, &Measure{DATE: splitKey[2], VALUE: data})
	}
	p, _ := json.Marshal(CaptorAndIata{IATA: iataCode, CAPTORNAME: captorName, VALUES: captorData})
	w.Write(p)
}

func GetDataByIataCodeAndCaptorAndYear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	iataCode := vars["iataCode"]
	captorName := vars["captorName"]
	year := vars["year"]

	var captorData []*Measure

	keys := bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/" + strings.ToUpper(captorName) + "/" + year + "-*")

	for _, key := range keys {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		captorData = append(captorData, &Measure{DATE: splitKey[2], VALUE: data})
	}
	p, _ := json.Marshal(CaptorAndIata{IATA: iataCode, CAPTORNAME: captorName, VALUES: captorData})
	w.Write(p)
}

func GetDataByIataCodeAndCaptorAndYearAndMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	iataCode := vars["iataCode"]
	captorName := vars["captorName"]
	year := vars["year"]
	month := vars["month"]

	var captorData []*Measure

	keys := bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/" + strings.ToUpper(captorName) + "/" + year + "-" + month + "-*")

	for _, key := range keys {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		captorData = append(captorData, &Measure{DATE: splitKey[2], VALUE: data})
	}
	p, _ := json.Marshal(CaptorAndIata{IATA: iataCode, CAPTORNAME: captorName, VALUES: captorData})
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

	var captorData []*Measure

	keys := bdd.GetAllKeyRegex(strings.ToUpper(iataCode) + "/" + strings.ToUpper(captorName) + "/" + year + "-" + month + "-" + day + "-*")

	for _, key := range keys {
		data := bdd.GetValue(key)
		splitKey := strings.Split(key, "/")
		captorData = append(captorData, &Measure{DATE: splitKey[2], VALUE: data})
	}
	p, _ := json.Marshal(CaptorAndIata{IATA: iataCode, CAPTORNAME: captorName, VALUES: captorData})
	w.Write(p)
}

func GetDataBetweenDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	captorName := vars["captorName"]
	start := vars["start"]
	end := vars["end"]

	var iataData []*Iata
	var mesure []*Measure

	startDate, _ := time.Parse("2006-01-02-15", start)
	endDate, _ := time.Parse("2006-01-02-15", end)

	var keys []string
	for d := startDate; d.After(endDate) == false; d = d.Add(time.Hour) {
		t := d.Format("2006-01-02-15")
		keys = append(keys, bdd.GetAllKeyRegex("*/"+strings.ToUpper(captorName)+"/"+t+"*")...)
	}

	mapIata := config.CODE_IATA

	for _, iata := range mapIata {
		mesure = nil
		for _, key := range keys {
			splitKey := strings.Split(key, "/")
			if splitKey[0] == iata {
				data := bdd.GetValue(key)
				mesure = append(mesure, &Measure{DATE: splitKey[2], VALUE: data})
			}
		}
		iataData = append(iataData, &Iata{IATA: iata, MEASURES: mesure})
	}

	p, _ := json.Marshal(BetweenDate{START: start, END: end, CAPTORNAME: captorName, VALUES: iataData})
	w.Write(p)
}

func GetAverageByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	date := vars["date"]

	pressureData := 0.0
	temperatureData := 0.0
	windData := 0.0

	dateDate, _ := time.Parse("2006-01-02", date)

	keysPressure := bdd.GetAllKeyRegex("*/PRESSURE/" + dateDate.Format("2006-01-02") + "*")
	keysTemperature := bdd.GetAllKeyRegex("*/TEMPERATURE/" + dateDate.Format("2006-01-02") + "*")
	keysWind := bdd.GetAllKeyRegex("*/WIND/" + dateDate.Format("2006-01-02") + "*")

	for _, key := range keysPressure {
		data := bdd.GetValue(key)
		chif, _ := strconv.ParseFloat(data, 64)
		pressureData += chif
	}
	for _, key := range keysTemperature {
		data := bdd.GetValue(key)
		chif, _ := strconv.ParseFloat(data, 64)
		temperatureData += chif
	}
	for _, key := range keysWind {
		data := bdd.GetValue(key)
		chif, _ := strconv.ParseFloat(data, 64)
		windData += chif
	}

	p := pressureData / float64(len(keysPressure))
	if math.IsNaN(p) {
		p = 0
	}
	t := temperatureData / float64(len(keysTemperature))
	if math.IsNaN(t) {
		t = 0
	}
	wi := windData / float64(len(keysWind))
	if math.IsNaN(wi) {
		wi = 0
	}

	j, _ := json.Marshal(DateAndAllCaptors{DATE: date, WIND: wi, TEMPERATURE: t, PRESSURE: p})
	w.Write(j)
}
