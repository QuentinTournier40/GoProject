package test

import (
	"context"
	"encoding/json"
	"goproject/api/router"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a router.App

func TestMain(m *testing.M) {
	// ...
	a.Initialize()
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	ctx := context.Background()
	if _, err := a.Client.Ping(ctx).Result(); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	ctx := context.Background()
	a.Client.FlushAll(ctx)
}

func TestEmptyTable(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/sensors", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetNonExistentSensor(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/sensor/11", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Sensor not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Sensor not found'. Got '%s'", m["error"])
	}
}

func TestGetSensor(t *testing.T) {
	clearTable()
	addSensors(1)
	req, _ := http.NewRequest("GET", "/sensor/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func addSensors(count int) {
	if count < 1 {
		count = 1
	}

	ctx := context.Background()
	for i := 0; i < count; i++ {
		a.Client.Set(ctx, "sensor:"+string(rune(i+1)), `{"id":"1","iata":"1","measuretype":"1","measurevalue":"1","measuredate":"1"}`, 0)
	}
}
