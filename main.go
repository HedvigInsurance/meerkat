package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/HedvigInsurance/meerkat/constants"
	"github.com/HedvigInsurance/meerkat/mappers"
	"github.com/HedvigInsurance/meerkat/queries"

	"github.com/gorilla/mux"
)

var euList mappers.SanctionEntities
var unList mappers.IndividualRoot

type response struct {
	Query  string `json:"query"`
	Result string `json:"result"`
}

func init() {
	log.Println("Meerkat started!")
	initialFetch := time.Now()
	euList = mappers.MapEuSanctionList()
	unList = mappers.MapUnSanctionList()
	log.Println("Both listed were initiated! It took: ", time.Since(initialFetch))
}

func main() {
	euListChannel := make(chan mappers.SanctionEntities)
	unListChannel := make(chan mappers.IndividualRoot)
	router := mux.NewRouter()

	go func() {
		for {
			time.Sleep(time.Hour)
			euListChannel <- mappers.MapEuSanctionList()
			log.Println("EU list was fetched")
		}
	}()

	go func() {
		for {
			time.Sleep(time.Hour)
			unListChannel <- mappers.MapUnSanctionList()
			log.Println("UN list was fetched")
		}
	}()

	go func() {
		router.HandleFunc("/api/check", func(w http.ResponseWriter, r *http.Request) {
			checkStatus(w, r)
		}).Methods(http.MethodGet).Queries("query", "{query}")
		log.Fatal(http.ListenAndServe(":80", router))
	}()

	for {
		select {
		case list := <-euListChannel:
			euList = list
			log.Println("EU list was updated")
		case list := <-unListChannel:
			unList = list
			log.Println("UN List was updated")
		}
	}
}

func checkStatus(w http.ResponseWriter, r *http.Request) {
	sanctionCheckStarted := time.Now()

	vars := mux.Vars(r)
	query := strings.Fields(vars["query"])

	log.Println("Sanctionlist search for", query, "Started")

	unResult := queries.QueryUnSanctionList(query, unList)

	if unResult == constants.FullHit {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response{vars["query"], unResult.ToString()})
		log.Println("UN Sanctioninst search for", query, "took", time.Since(sanctionCheckStarted), "Result:", unResult.ToString())
	} else {
		euResult := queries.QueryEuSanctionList(query, euList)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response{vars["query"], euResult.ToString()})
		log.Println("EU Sanctionlist search for", query, "took", time.Since(sanctionCheckStarted), "Result:", euResult.ToString())
	}
}
