package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func wait(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	waitTime := vars["time"]

	i, err := strconv.ParseInt(waitTime, 10, 64)
	if err != nil {
		http.Error(w, `{"error": "Bad Request"}`, http.StatusBadRequest)
		return
	}

	time.Sleep(time.Duration(i) * time.Second)
	fmt.Fprintf(w, `{"message": "bow-wow!"}`)

}

func waitRandom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rand.Seed(time.Now().UnixNano())

	ra := rand.Int63n(10) // 1 ~ 10
	time.Sleep(time.Duration(ra) * time.Second)

	fmt.Fprintf(w, `{"message": "bow-wow!"}`)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"health": "ok"}`)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()
	router.Path("/").HandlerFunc(index)
	router.Path("/wait/random").HandlerFunc(waitRandom)
	router.Path("/wait/{time}").HandlerFunc(wait)

	h1 := handlers.CombinedLoggingHandler(os.Stdout, router)

	http.ListenAndServe(":"+port, h1)
}
