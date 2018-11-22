package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/HedvigInsurance/meerkat/mappers"
	"github.com/HedvigInsurance/meerkat/queries"

	"github.com/gorilla/mux"
)

var mu sync.RWMutex
var euList mappers.SanctionEntites
var unList mappers.IndividualRoot

type Response struct {
	Query  string
	Result string
}

func main() {
	start_fetch := time.Now()
	euList = mappers.MapEuSanctionList()
	unList = mappers.MapUnSanctionList()

	log.Println("Fetching took ", time.Since(start_fetch))

	go func() {
		for {
			time.Sleep(240)

			euTemp := mappers.MapEuSanctionList()
			mu.Lock()
			euList = euTemp
			mu.Unlock()
		}
	}()

	go func() {
		for {
			time.Sleep(240)
			unTemp := mappers.MapUnSanctionList()
			mu.Lock()
			unList = unTemp
			mu.Unlock()
		}
	}()

	router := mux.NewRouter()
	router.HandleFunc("/api/check", checkStatus).Methods(http.MethodGet).Queries("query", "{query}")
	router.HandleFunc("/api/check", checkStatusWithFirstAndLastName).Methods(http.MethodGet).Queries("firstName", "{firstName}", "lastName", "{lastName}")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func checkStatus(w http.ResponseWriter, r *http.Request) {
	start_sanct := time.Now()

	vars := mux.Vars(r)
	query := strings.Split(vars["query"], " ")

	mu.Lock()
	result := queries.QueryEUsanctionList(query, euList)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{vars["query"], result.ToString()})
	log.Println("Sanctionlist took ", time.Since(start_sanct))
}

func checkStatusWithFirstAndLastName(w http.ResponseWriter, r *http.Request) {
	start_sanct := time.Now()

	vars := mux.Vars(r)
	query := []string{vars["firstName"], vars["lastName"]}

	mu.Lock()
	result := queries.QueryEUsanctionList(query, euList)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{vars["firstName"] + " " + vars["lastName"], result.ToString()})
	log.Println("Sanctionlist took ", time.Since(start_sanct))
}
