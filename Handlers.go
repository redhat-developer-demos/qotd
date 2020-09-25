package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AllQuotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(quotes)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "qotd")
}

func OneQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["quoteId"])
	if err == nil {
		json.NewEncoder(w).Encode(quotes[id])
	}
}

func RandomQuote(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(quotes[rand.Intn(6)])
}

func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "2.0.0")
}

func WrittenIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go")
}
