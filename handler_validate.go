package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	bannedWords := []string{"kerfuffle", "sharbert", "fornax"}
	cleanedChirp := []string{}
	chirpWords := strings.Split(params.Body, " ")
	for _, word := range chirpWords {
		for _, banned := range bannedWords {
			if strings.ToLower(word) == banned {
				word = "****"
			}
		}
		cleanedChirp = append(cleanedChirp, word)
	}

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: strings.Join(cleanedChirp, " "),
	})
}