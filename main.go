package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ToRead struct {
	Title    string
	Referrer string
}

type ToReads []ToRead

func getToReadsHandler(w http.ResponseWriter, r *http.Request) {
	toReads := ToReads{
		ToRead{Title: "计算广告"},
		ToRead{Title: "智能商业"},
	}

	json.NewEncoder(w).Encode(toReads)
}

func main() {
	http.HandleFunc("/to_read/", getToReadsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
