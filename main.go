package main

import (
	"log"
	"net/http"
)

func handleCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusBadRequest, "invalid method")
		return
	}

	card := returnCard(r)

	computedCD, realCD := returnCDs(card)

	if realCD == computedCD {
		respondWithType(w, 200, msgResponse{}, "valid card number")
	} else {
		respondWithError(w, http.StatusBadRequest, "invalid card number")
	}
}

func handleCheckAndSuggest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusBadRequest, "invalid method")
		return
	}
	
	card := returnCard(r)

	computedCD, realCD := returnCDs(card)

	provider, isIndustry := providerCheck(card.Number)

	if realCD == computedCD && !isIndustry {
		respondWithType(w, 200, msgResponse{}, provider)
	} else if realCD == computedCD {
		respondWithType(w, 200, industryResponse{}, provider)
	} else {
		respondWithError(w, http.StatusBadRequest, "invalid card number")
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/check", handleCheck)
	mux.HandleFunc("/check_suggest", handleCheckAndSuggest)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
