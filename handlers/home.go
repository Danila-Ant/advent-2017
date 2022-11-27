package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// home returns a simple HTTP handler function which writes a response.
func home(buildTime, commit, release string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		info := struct {
			BuildTime string `json:"buildTime1"`
			Commit    string `json:"commit1"`
			Release   string `json:"release1"`
		}{
			buildTime, commit, release,
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
}
