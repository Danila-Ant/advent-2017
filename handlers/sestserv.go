package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func sestserv(w http.ResponseWriter, _ *http.Request) {

	info := struct {
		BuildTime string `json:"buildTime"`
		Commit    string `json:"commit"`
		Release   string `json:"release"`
		Pole4     string `json:"pole4"`
	}{
		"kol1", "kol2", "kol3", "kol4",
	}

	body, err := json.Marshal(info)
	if err != nil {
		log.Printf("Could not encode info data: %v", err)
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)

}
