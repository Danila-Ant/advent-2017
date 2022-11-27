package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func sestserv(w http.ResponseWriter, _ *http.Request) {

	var resserv = "test"

	body, err := json.Marshal(resserv)
	if err != nil {
		log.Printf("Could not encode info data: %v", err)
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}
