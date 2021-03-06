package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"

	"github.com/HedvigInsurance/meerkat/constants"
	"github.com/HedvigInsurance/meerkat/mappers"
	"github.com/HedvigInsurance/meerkat/queries"

	"github.com/gorilla/mux"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var euList mappers.SanctionEntities
var unList mappers.IndividualRoot

type response struct {
	Query  string `json:"query"`
	Result string `json:"result"`
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Meerkat started!")
	initialFetch := time.Now()
	euList = mappers.MapEuSanctionList()
	unList = mappers.MapUnSanctionList()
	log.Info("Both listed were initiated! It took: ", time.Since(initialFetch))
}

func main() {
	//Starts Datadog tracer
	tracer.Start()
	defer tracer.Stop()
	// Create a traced mux router.
	mux := muxtrace.NewRouter(muxtrace.WithServiceName("meerkat"))

	euListChannel := make(chan mappers.SanctionEntities)
	unListChannel := make(chan mappers.IndividualRoot)

	go func() {
		for {
			time.Sleep(time.Hour)
			euListChannel <- mappers.MapEuSanctionList()
			log.Info("EU list was fetched")
		}
	}()

	go func() {
		for {
			time.Sleep(time.Hour)
			unListChannel <- mappers.MapUnSanctionList()
			log.Info("UN list was fetched")
		}
	}()

	go func() {
		mux.HandleFunc("/api/check", func(w http.ResponseWriter, r *http.Request) {
			checkStatus(w, r)
		}).Methods(http.MethodGet).Queries("query", "{query}")
		log.Fatal(http.ListenAndServe(":80", mux))
	}()

	for {
		select {
		case list := <-euListChannel:
			euList = list
			log.Info("EU list was updated")
		case list := <-unListChannel:
			unList = list
			log.Info("UN List was updated")
		}
	}
}

func checkStatus(w http.ResponseWriter, r *http.Request) {
	sanctionCheckStarted := time.Now()

	vars := mux.Vars(r)
	query := strings.Fields(vars["query"])

	unResult := queries.QueryUnSanctionList(query, unList)

	if unResult == constants.FullHit {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response{vars["query"], unResult.ToString()})
		log.Info("UN Sanctioninst search completed, took", time.Since(sanctionCheckStarted), "Result:", unResult.ToString())
	} else {
		euResult := queries.QueryEuSanctionList(query, euList)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response{vars["query"], euResult.ToString()})
		log.Info("EU Sanctionlist search completed, took", time.Since(sanctionCheckStarted), "Result:", euResult.ToString())
	}
}
