package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/HedvigInsurance/meerkat/constants"
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

	log.Println("Initial fetching took ", time.Since(start_fetch))

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
	log.Fatal(http.ListenAndServe(":80", router))
}

func checkStatus(w http.ResponseWriter, r *http.Request) {
	start_sanct := time.Now()

	vars := mux.Vars(r)
	query := strings.Split(vars["query"], " ")

	mu.Lock()
	euResult := queries.QueryEUsanctionList(query, euList)
	mu.Unlock()

	if euResult == constants.FullHit {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{vars["query"], euResult.ToString()})
		log.Println("EU Sanctioninst search for", query, "took", time.Since(start_sanct), "Result:", euResult.ToString())
	} else {
		mu.Lock()
		unResult := queries.QueryUNsanctionList(query, unList)
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{vars["query"], unResult.ToString()})
		log.Println("UN Sanctionlist search for", query, "took", time.Since(start_sanct), "Result:", unResult.ToString())
	}
}
