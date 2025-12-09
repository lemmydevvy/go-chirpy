package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/lemmydevvy/go-chirpy/internal/database"
)

func (cfg *apiConfig) handlerChirpsList(w http.ResponseWriter, r *http.Request) {
	authorId := r.URL.Query().Get("author_id")
	var chirps []database.Chirp
	var err error

	if authorId != "" {
		id, _ := uuid.Parse(authorId)
		chirps, err = cfg.db.GetChirpsByUserID(r.Context(), id)
	} else {
		chirps, err = cfg.db.GetChirps(r.Context())
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
		return
	}

	chirpsList := []Chirp{}
	for _, chirp := range chirps {
		chirpsList = append(chirpsList, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			UserID:    chirp.UserID,
			Body:      chirp.Body,
		})
	}

	respondWithJSON(w, http.StatusOK, chirpsList)
}