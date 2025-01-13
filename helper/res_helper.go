package helper

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/devdhanadiya/rss-aggregator/model"
)

func ResJson(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		w.WriteHeader(statusCode)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func ResError(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode > 499 {
		log.Println("Responding with 5XX error", msg)
	}
	ResJson(w, statusCode, model.ErrResponse{Error: msg})
}
