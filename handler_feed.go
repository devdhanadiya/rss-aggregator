package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/devdhanadiya/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type paramaters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := paramaters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf(" parsing error: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("could not create user error: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedToFeed(feed))
}
