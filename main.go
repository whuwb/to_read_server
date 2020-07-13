package main

import (
	"ToRead/db"
	"encoding/json"
	"log"
	"net/http"
)

func getToReadsHandler(w http.ResponseWriter, r *http.Request) {
	toReads := db.FetchToReads()

	json.NewEncoder(w).Encode(toReads)
}

func main() {
	http.HandleFunc("/to_read/", getToReadsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
