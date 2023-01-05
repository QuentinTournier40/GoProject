package captorService

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"goproject/internal/bdd"
	"goproject/internal/config"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ------------------------------ TYPES JSON ------------------------------

type AllCaptors struct {
	IATA        string     `json:"iata" example:"CDG" description:"Identifiant d'un aeroport"`
	PRESSURE    []*Measure `json:"pressure" example:"[{\"date\":\"2022-12-25-12-00-00\", \"value\":\"1015.13\"}]" description:"Liste des valeurs pour les capteurs de pression"`
	TEMPERATURE []*Measure `json:"temperature" example:"[{\"date\":\"2022-12-25-12-00-00\", \"value\":\"25.5\"}]" description:"Liste des valeurs pour les capteurs de temperature"`
	WIND        []*Measure `json:"wind" example:"[{\"date\":\"2022-12-25-12-00-00\", \"value\":\"60.6\"}]" description:"Liste des valeurs pour les capteurs de vitesse du vent"`
}

type Measure struct {
	DATE  string `json:"date" example:"2022-12-25-12-00-00" description:"Date (YYYY-MM-DD-HH-MM-SS) de l'enregistrement de la valeur"`
	VALUE string `json:"value" example:"25.8" description:"Valeur enregistré par le capteur"`
}

type Captor struct {
	CAPTORNAME string  `json:"captorName" example:"pressure" description:"Nom du type du capteur"`
	VALUES     []*Iata `json:"values" example:"[{\"iata\": \"CDG\", \"measures\":[{\"date\":\"2022-12-25-12-00-00\", \"value\":\"60.6\"}]}]" description:"Liste de code iata et de ses valeurs"`
}

type Iata struct {
	IATA     string     `json:"iata" example:"CDG" description:"Identifiant d'un aeroport"`
	MEASURES []*Measure `json:"measures" example:"[{\"date\":\"2022-12-25-12-00-00\", \"value\":\"60.6\"}]" description:"Liste de valeurs d'un capteur"`
}

type CaptorAndIata struct {
	IATA       string     `json:"iata" example:"CDG" description:"Identifiant d'un aeroport"`
	CAPTORNAME string     `json:"captorName" example:"pressure" description:"Nom du type du capteur"`
	VALUES     []*Measure `json:"values" example:"[{\"date\":\"2022-12-25-12-00-00\", \"value\":\"60.6\"}]" description:"Liste de valeurs associé a un code iata et un type de capteur"`
}

type BetweenDate struct {
	START      string  `json:"start" example:"2022-12-25-12" description:"Date (YYYY-MM-DD-HH) du début de la plage horaire"`
	END        string  `json:"end" example:"2022-12-25-13" description:"Date (YYYY-MM-DD-HH) de la fin de la plage horaire"`
	CAPTORNAME string  `json:"captorName" example:"pressure" description:"Nom du type du capteur"`
	VALUES     []*Iata `json:"values" example:"[{\"date\":\"2022-12-25-12-00-00\", \"value\":\"60.6\"}]" description:"Liste de valeurs associé a un code iata et un type de capteur compris dans le plage horaire"`
}

type DateAndAllCaptors struct {
	DATE        string  `json:"date" example:"2022-12-25" description:"Date (YYYY-MM-DD) du jour ou l'on veut connaitre les moyennes des valeurs"`
	PRESSURE    float64 `json:"pressure" example:"950.12" description:"Valeur moyenne des données de pression"`
	TEMPERATURE float64 `json:"temperature" example:"25.3" description:"Valeur moyenne des données de temperature"`
	WIND        float64 `json:"wind" example:"60.2" description:"Valeur moyenne des données de votesse du vent"`
}

// ------------------------------ SERVICES ------------------------------

// @Title Get all data by iata code
// @Description Obtenir tous les relevés de mesure selon un code iata
// @Param iataCode path string true "Code iata"
// @Success 200 {object} AllCaptors "AllCaptors JSON"
// @Route /get/data-by-iata-code/{iataCode} [get]
func GetDataByIataCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	iataCode := strings.ToUpper(vars["iataCode"])

	var pressureMeasures []*Measure
	var temperatureMeasures []*Measure
	var windMeasures []*Measure

	dataPressure := bdd.GetValuesBetween2Index(iataCode+"/PRESSURE", 0, -1)
	dataTemperature := bdd.GetValuesBetween2Index(iataCode+"/TEMPERATURE", 0, -1)
	dataWind := bdd.GetValuesBetween2Index(iataCode+"/WIND", 0, -1)

	for _, value := range dataPressure {
		splitValue := strings.Split(value, ":")
		dateUnixInt, _ := strconv.ParseInt(splitValue[0], 10, 64)
		date := time.Unix(dateUnixInt, 0)
		pressureMeasures = append(pressureMeasures, &Measure{DATE: date.Format("2006-01-02-15-04-05"), VALUE: splitValue[1]})
	}
	for _, value := range dataTemperature {
		splitValue := strings.Split(value, ":")
		dateUnixInt, _ := strconv.ParseInt(splitValue[0], 10, 64)
		date := time.Unix(dateUnixInt, 0)
		temperatureMeasures = append(temperatureMeasures, &Measure{DATE: date.Format("2006-01-02-15-04-05"), VALUE: splitValue[1]})
	}
	for _, value := range dataWind {
		splitValue := strings.Split(value, ":")
		dateUnixInt, _ := strconv.ParseInt(splitValue[0], 10, 64)
		date := time.Unix(dateUnixInt, 0)
		windMeasures = append(windMeasures, &Measure{DATE: date.Format("2006-01-02-15-04-05"), VALUE: splitValue[1]})
	}

	p, _ := json.Marshal(AllCaptors{IATA: iataCode, PRESSURE: pressureMeasures, TEMPERATURE: temperatureMeasures, WIND: windMeasures})
	w.Write(p)
}

// @Title Get amount of data by iata code
// @Description Obtenir un nombre donnée de relevés de mesure selon un code iata
// @Param iataCode path string true "Code iata"
// @Param number path string true "Nombre"
// @Success 200 {object} AllCaptors "AllCaptors JSON"
// @Route /get/data-by-iata-code-and-number/{iataCode}/{number} [get]
func GetDataByIataCodeForXData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	iataCode := strings.ToUpper(vars["iataCode"])
	number, _ := strconv.ParseInt(vars["number"], 10, 64)

	var pressureMeasures []*Measure
	var temperatureMeasures []*Measure
	var windMeasures []*Measure

	dataPressure := bdd.GetValuesBetween2Index(iataCode+"/PRESSURE", -number, -1)
	dataTemperature := bdd.GetValuesBetween2Index(iataCode+"/TEMPERATURE", -number, -1)
	dataWind := bdd.GetValuesBetween2Index(iataCode+"/WIND", -number, -1)

	for _, value := range dataPressure {
		splitValue := strings.Split(value, ":")
		dateUnixInt, _ := strconv.ParseInt(splitValue[0], 10, 64)
		date := time.Unix(dateUnixInt, 0)
		pressureMeasures = append(pressureMeasures, &Measure{DATE: date.Format("2006-01-02-15-04-05"), VALUE: splitValue[1]})
	}
	for _, value := range dataTemperature {
		splitValue := strings.Split(value, ":")
		dateUnixInt, _ := strconv.ParseInt(splitValue[0], 10, 64)
		date := time.Unix(dateUnixInt, 0)
		temperatureMeasures = append(temperatureMeasures, &Measure{DATE: date.Format("2006-01-02-15-04-05"), VALUE: splitValue[1]})
	}
	for _, value := range dataWind {
		splitValue := strings.Split(value, ":")
		dateUnixInt, _ := strconv.ParseInt(splitValue[0], 10, 64)
		date := time.Unix(dateUnixInt, 0)
		windMeasures = append(windMeasures, &Measure{DATE: date.Format("2006-01-02-15-04-05"), VALUE: splitValue[1]})
	}

	p, _ := json.Marshal(AllCaptors{IATA: iataCode, PRESSURE: pressureMeasures, TEMPERATURE: temperatureMeasures, WIND: windMeasures})
	w.Write(p)
}

// @Title Get all data by captor
// @Description Obtenir tous les relevés de mesure d'un type de capteur
// @Param captorName path string true "Captor name"
// @Success 200 {object} Captor "Captor JSON"
// @Route /get/data-by-captorName/{captorName} [get]
func GetDataByCaptor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

// @Title Get data by iata code and captor name
// @Description Obtenir tous les relevés de mesure d'un aeroport et d un seul type de capteur.
// @Param captorName path string true "Captor name"
// @Param iataCode path string true "Code iata"
// @Success 200 {object} CaptorAndIata "CaptorAndIata JSON"
// @Route /get/data/{iataCode}/{captorName} [get]
func GetDataByIataCodeAndCaptor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

// @Title Get data by iata code, captor name and year
// @Description Obtenir tous les relevés de mesure d'un aeroport et d un seul type de capteur d une année precise.
// @Param captorName path string true "Captor name"
// @Param iataCode path string true "Code iata"
// @Param year path string true "Year"
// @Success 200 {object} CaptorAndIata "CaptorAndIata JSON"
// @Route /get/data/{iataCode}/{captorName}/{year} [get]
func GetDataByIataCodeAndCaptorAndYear(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

// @Title Get data by iata code, captor name, year and month
// @Description Obtenir tous les relevés de mesure d'un aeroport et d'un seul type de capteur d une année et d un mois precis.
// @Param captorName path string true "Captor name"
// @Param iataCode path string true "Code iata"
// @Param year path string true "Year"
// @Param month path string true "Month"
// @Success 200 {object} CaptorAndIata "CaptorAndIata JSON"
// @Route /get/data/{iataCode}/{captorName}/{year}/{month} [get]
func GetDataByIataCodeAndCaptorAndYearAndMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

// @Title Get data by iata code, captor name, year, month and day
// @Description Obtenir tous les relevés de mesure d un aeroport et d un seul type de capteur d une année, d un mois et d un jour precis.
// @Param captorName path string true "Captor name"
// @Param iataCode path string true "Code iata"
// @Param year path string true "Year"
// @Param month path string true "Month"
// @Param day path string true "Day"
// @Success 200 {object} CaptorAndIata "CaptorAndIata JSON"
// @Route /get/data/{iataCode}/{captorName}/{year}/{month}/{day} [get]
func GetDataByIataCodeAndCaptorAndYearAndMonthAndDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

// @Title Get data between 2 dates
// @Description Obtenir tous les relevés de mesure d un type de capteur dans une plage de temps donnée.
// @Param captorName path string true "Captor name"
// @Param start path string true "Start"
// @Param end path string true "End"
// @Success 200 {object} BetweenDate "CaptorAndIata JSON"
// @Route /get/data-between-dates/{captorName}/{start}/{end} [get]
func GetDataBetweenDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	captorName := strings.ToUpper(vars["captorName"])
	start := vars["start"]
	end := vars["end"]

	var iataData []*Iata
	var measures []*Measure
	mapIata := config.CODE_IATA

	startDate, _ := time.Parse("2006-01-02-15", start)
	endDate, _ := time.Parse("2006-01-02-15", end)

	startDateUnix := startDate.Unix()
	endDateUnix := endDate.Unix()

	for _, iata := range mapIata {
		measures = nil
		for _, value := range bdd.GetValuesBetween2Score(iata+"/"+captorName, startDateUnix, endDateUnix) {
			splitValue := strings.Split(value, ":")
			dateUnixInt, _ := strconv.ParseInt(splitValue[0], 10, 64)
			date := time.Unix(dateUnixInt, 0)
			measures = append(measures, &Measure{DATE: date.Format("2006-01-02-15-04-05"), VALUE: splitValue[1]})
		}
		iataData = append(iataData, &Iata{IATA: iata, MEASURES: measures})
	}

	p, _ := json.Marshal(BetweenDate{START: start, END: end, CAPTORNAME: captorName, VALUES: iataData})
	w.Write(p)
}

// @Title Get average data from all captor on specific day
// @Description Obtenir la moyenne de tous les releves d'un jour donné.
// @Param date path string true "Date"
// @Success 200 {object} DateAndAllCaptors "CaptorAndIata JSON"
// @Route /get/average-data/{date} [get]
func GetAverageByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	date := vars["date"]

	pressureSum := 0.0
	temperatureSum := 0.0
	windSum := 0.0

	cpt1 := 0
	cpt2 := 0
	cpt3 := 0

	dateDate, _ := time.Parse("2006-01-02", date)

	startDateUnix := dateDate.Unix()
	endDateUnix := dateDate.Add(time.Hour * 24).Unix()

	for _, key := range bdd.GetAllKeyRegex("*/PRESSURE") {
		for _, value := range bdd.GetValuesBetween2Score(key, startDateUnix, endDateUnix) {
			fmt.Println('a')
			splitValue := strings.Split(value, ":")
			chiffre, _ := strconv.ParseFloat(splitValue[1], 64)
			pressureSum += chiffre
			cpt1 += 1
		}
	}
	for _, key := range bdd.GetAllKeyRegex("*/TEMPERATURE") {
		for _, value := range bdd.GetValuesBetween2Score(key, startDateUnix, endDateUnix) {
			splitValue := strings.Split(value, ":")
			chiffre, _ := strconv.ParseFloat(splitValue[1], 64)
			temperatureSum += chiffre
			cpt2 += 1
		}
	}
	for _, key := range bdd.GetAllKeyRegex("*/WIND") {
		for _, value := range bdd.GetValuesBetween2Score(key, startDateUnix, endDateUnix) {
			splitValue := strings.Split(value, ":")
			chiffre, _ := strconv.ParseFloat(splitValue[1], 64)
			windSum += chiffre
			cpt3 += 1
		}
	}

	averagePressure := pressureSum / float64(cpt1)
	averageTemperature := temperatureSum / float64(cpt2)
	averageWind := windSum / float64(cpt3)
	//
	if math.IsNaN(averagePressure) {
		averagePressure = math.SmallestNonzeroFloat64
	}
	if math.IsNaN(averageTemperature) {
		averageTemperature = math.SmallestNonzeroFloat64
	}
	if math.IsNaN(averageWind) {
		averageWind = math.SmallestNonzeroFloat64
	}

	fmt.Println(averagePressure)

	j, _ := json.Marshal(DateAndAllCaptors{DATE: date, WIND: averageWind, TEMPERATURE: averageTemperature, PRESSURE: averagePressure})
	w.Write(j)
}
