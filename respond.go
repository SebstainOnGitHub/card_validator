package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type errResponse struct {
	Error interface{} `json:"error"`
}

type msgResponse struct {
	Message string `json:"message"`
}

type providerResponse struct {
	Provider string `json:"provider"`
}

type industryResponse struct {
	Industry string `json:"industry"`
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON response %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, payload interface{}) {
	if code > 499 {
		log.Println("Responding with 5XX error:", payload)
	}

	respondWithJSON(w, code, errResponse{
		Error: payload,
	})
}

func respondWithType(w http.ResponseWriter, code int, typeMessage interface{}, msg string) {
	if code >= 200 && code < 300 {
		log.Println("Responding with 2XX code:", msg)
	}

	switch typeMessage.(type) {
	case providerResponse:
		respondWithJSON(w, code, providerResponse{
			Provider: msg,
		})
	case msgResponse:
		respondWithJSON(w, code, msgResponse{
			Message: msg,
		})
	case industryResponse:
		respondWithJSON(w, code, industryResponse{
			Industry: msg,
		})
	}
}