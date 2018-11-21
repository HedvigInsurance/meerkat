package main

import (
	"log"
	"meerkat/mappers"
	"meerkat/queries"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type SanctionResult uint16

var mu sync.RWMutex
var euList mappers.SanctionEntites
var unList mappers.IndividualRoot

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

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/check", checkSanctionStatus).Methods(http.MethodGet).Queries("query", "{result}")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func checkSanctionStatus(w http.ResponseWriter, r *http.Request) {

	start_sanct := time.Now()

	vars := mux.Vars(r)
	query := strings.Split(vars["result"], ",")

	mu.Lock()
	result := queries.QueryEUsanctionList(query, euList)
	mu.Unlock()

	w.WriteHeader(200)
	w.Write([]byte("TEST"))
	log.Println("Sanction Result ", result)
	log.Println("Sanctionlist took ", time.Since(start_sanct))
}
