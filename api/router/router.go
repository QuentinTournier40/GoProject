package router

import (
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
	"goproject/api/captor"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	Router *mux.Router
	Client *redis.Client
}

func (a *App) Initialize() {
	a.Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) getSensor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	s := model.Sensor{ID: id}
	if err := s.GetSensor(a.Client); err != nil {
		switch err {
		case redis.Nil:
			respondWithError(w, http.StatusNotFound, "Sensor not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, s)
}

func (a *App) getSensors(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))
	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}
	sensors, err := model.GetSensors(a.Client)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, sensors)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/sensors", a.getSensors).Methods("GET")
	a.Router.HandleFunc("/sensor/{id:[0-9]+}", a.getSensor).Methods("GET")
}
