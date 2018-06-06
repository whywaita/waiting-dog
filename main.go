package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func wait(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	waitTime := vars["time"]

	i, err := strconv.ParseInt(waitTime, 10, 64)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	time.Sleep(time.Duration(i) * time.Second)
	fmt.Fprintf(w, "bow-wow!")

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"health": "ok"}`)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()
	router.Path("/").HandlerFunc(index)
	router.Path("/wait/{time}").HandlerFunc(wait)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("[ERROR]", err)
	}
}
