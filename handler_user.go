package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/devdhanadiya/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *ApiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type paramaters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := paramaters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf(" parsing error: %v", err))
		return
	}

	newUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("could not create user error: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUsertoUser(newUser))
}

func (apiCfg *ApiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUsertoUser(user))
}

func (apiCfg *ApiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could'nt  create user: %v", err))
		return
	}

	respondWithJSON(w, 200, databasePostsToPosts(posts))
}
