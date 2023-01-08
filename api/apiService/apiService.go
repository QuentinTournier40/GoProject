package apiService

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goproject/internal/bdd"
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
	WIND        float64 `json:"wind" example:"60.2" description:"Valeur moyenne des données de vitesse du vent"`
}

// ------------------------------ SERVICES ------------------------------

// @Title GetDataByIataCode
// @Description Obtenir tous les relevés de mesure selon un code iata
// @Param code path string true "Code iata"
// @Success 200 {object} AllCaptors "AllCaptors JSON"
// @Route /iata/{code} [get]
func GetDataByIataCode(w http.ResponseWriter, r *http.Request) {
	setHeader(w)

	iataCode := strings.ToUpper(mux.Vars(r)["code"])

	var pressureMeasures []*Measure
	var temperatureMeasures []*Measure
	var windMeasures []*Measure

	dataPressure := bdd.GetValuesBetween2Index(iataCode+"/PRESSURE", 0, -1)
	dataTemperature := bdd.GetValuesBetween2Index(iataCode+"/TEMPERATURE", 0, -1)
	dataWind := bdd.GetValuesBetween2Index(iataCode+"/WIND", 0, -1)

	pressureMeasures = createMeasure(dataPressure)
	temperatureMeasures = createMeasure(dataTemperature)
	windMeasures = createMeasure(dataWind)

	p, _ := json.Marshal(AllCaptors{IATA: iataCode, PRESSURE: pressureMeasures, TEMPERATURE: temperatureMeasures, WIND: windMeasures})
	w.Write(p)
}

// @Title GetDataByIataCodeForXData
// @Description Obtenir un nombre donné de relevés de mesure selon un code iata
// @Param code path string true "Code iata"
// @Param number path string true "Nombre"
// @Success 200 {object} AllCaptors "AllCaptors JSON"
// @Route /iata/{code}/number/{number} [get]
func GetDataByIataCodeForXData(w http.ResponseWriter, r *http.Request) {
	setHeader(w)

	iataCode := strings.ToUpper(mux.Vars(r)["code"])
	number, _ := strconv.ParseInt(mux.Vars(r)["number"], 10, 64)

	var pressureMeasures []*Measure
	var temperatureMeasures []*Measure
	var windMeasures []*Measure

	dataPressure := bdd.GetValuesBetween2Index(iataCode+"/PRESSURE", -number, -1)
	dataTemperature := bdd.GetValuesBetween2Index(iataCode+"/TEMPERATURE", -number, -1)
	dataWind := bdd.GetValuesBetween2Index(iataCode+"/WIND", -number, -1)

	pressureMeasures = createMeasure(dataPressure)
	temperatureMeasures = createMeasure(dataTemperature)
	windMeasures = createMeasure(dataWind)

	p, _ := json.Marshal(AllCaptors{IATA: iataCode, PRESSURE: pressureMeasures, TEMPERATURE: temperatureMeasures, WIND: windMeasures})
	w.Write(p)
}

// @Title GetDataByCaptor
// @Description Obtenir tous les relevés de mesure d'un type de capteur
// @Param sensorName path string true "Captor name"
// @Success 200 {object} Captor "Captor JSON"
// @Route /sensor/{sensorName} [get]
func GetDataByCaptor(w http.ResponseWriter, r *http.Request) {
	setHeader(w)

	sensorName := strings.ToUpper(mux.Vars(r)["sensorName"])

	var iataData []*Iata
	var measures []*Measure

	tabKeys := bdd.GetAllKeyRegex("*/" + sensorName)

	for _, key := range tabKeys {
		iata := strings.Split(key, "/")[0]
		measures = createMeasure(bdd.GetValuesBetween2Index(key, 0, -1))
		iataData = append(iataData, &Iata{IATA: iata, MEASURES: measures})
		measures = nil
	}

	p, _ := json.Marshal(Captor{CAPTORNAME: sensorName, VALUES: iataData})
	w.Write(p)
}

// @Title GetDataByIataCodeAndCaptor
// @Description Obtenir tous les relevés de mesure d'un aeroport et d un seul type de capteur.
// @Param code path string true "Code iata"
// @Param sensorName path string true "Captor name"
// @Success 200 {object} CaptorAndIata "CaptorAndIata JSON"
// @Route /iata/{code}/sensor/{sensorName} [get]
func GetDataByIataCodeAndCaptor(w http.ResponseWriter, r *http.Request) {
	setHeader(w)

	iataCode := strings.ToUpper(mux.Vars(r)["code"])
	sensorName := strings.ToUpper(mux.Vars(r)["sensorName"])

	var measures []*Measure

	data := bdd.GetValuesBetween2Index(iataCode+"/"+sensorName, 0, -1)

	measures = createMeasure(data)

	p, _ := json.Marshal(CaptorAndIata{IATA: iataCode, CAPTORNAME: sensorName, VALUES: measures})
	w.Write(p)
}

// @Title GetDataBetweenDates
// @Description Obtenir tous les relevés de mesure d un type de capteur dans une plage de temps donnée.
// @Param sensorName path string true "Captor name"
// @Param start_date path string true "Start"
// @Param end_date path string true "End"
// @Success 200 {object} BetweenDate "CaptorAndIata JSON"
// @Route /sensor/{sensorName}/between/{start_date}/to/{end_date} [get]
func GetDataBetweenDates(w http.ResponseWriter, r *http.Request) {
	setHeader(w)

	sensorName := strings.ToUpper(mux.Vars(r)["sensorName"])
	startDate, _ := time.Parse("2006-01-02-15", mux.Vars(r)["start_date"])
	endDate, _ := time.Parse("2006-01-02-15", mux.Vars(r)["end_date"])

	var iataData []*Iata
	var measures []*Measure

	tabIata := bdd.GetAllKeyRegex("*/" + sensorName)

	for _, key := range tabIata {
		iata := strings.Split(key, "/")[0]
		measures = createMeasure(bdd.GetValuesBetween2Score(key, startDate.Unix(), endDate.Unix()))
		iataData = append(iataData, &Iata{IATA: iata, MEASURES: measures})
		measures = nil
	}

	p, _ := json.Marshal(BetweenDate{START: mux.Vars(r)["start_date"], END: mux.Vars(r)["end_date"], CAPTORNAME: sensorName, VALUES: iataData})
	w.Write(p)
}

// @Title GetAverageByDate
// @Description Obtenir la moyenne de tous les relevés d'un jour donné.
// @Param date path string true "Date"
// @Success 200 {object} DateAndAllCaptors "CaptorAndIata JSON"
// @Route /averages/{date} [get]
func GetAverageByDate(w http.ResponseWriter, r *http.Request) {
	setHeader(w)

	date, _ := time.Parse("2006-01-02", mux.Vars(r)["date"])

	startDateUnix := date.Unix()
	endDateUnix := date.Add(time.Hour * 24).Unix()

	averagePressure := getAverage("PRESSURE", startDateUnix, endDateUnix)
	averageTemperature := getAverage("TEMPERATURE", startDateUnix, endDateUnix)
	averageWind := getAverage("WIND", startDateUnix, endDateUnix)

	j, _ := json.Marshal(DateAndAllCaptors{DATE: mux.Vars(r)["date"], WIND: averageWind, TEMPERATURE: averageTemperature, PRESSURE: averagePressure})
	w.Write(j)
}

// ------------------------------------------------------------------------------------------------------

func createMeasure(data []string) []*Measure {
	var measures []*Measure
	for _, value := range data {
		splitValue := strings.Split(value, ":")
		measures = append(measures, &Measure{DATE: splitValue[0], VALUE: splitValue[1]})
	}
	return measures
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func getAverage(sensorName string, startUnix, endUnix int64) float64 {
	sum, average := 0.0, math.SmallestNonzeroFloat64
	cpt := 0
	for _, key := range bdd.GetAllKeyRegex("*/" + sensorName) {
		for _, value := range bdd.GetValuesBetween2Score(key, startUnix, endUnix) {
			splitValue := strings.Split(value, ":")
			number, _ := strconv.ParseFloat(splitValue[1], 64)
			sum += number
			cpt++
		}
	}
	if !math.IsNaN(sum / float64(cpt)) {
		average = sum / float64(cpt)
	}
	return average
}
