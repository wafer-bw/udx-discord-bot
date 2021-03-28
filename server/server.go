package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/wafer-bw/udx-discord-bot/api"
)

func mux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api", apiHandler)
	return mux
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	api.Handler(w, r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func getEnv(key string, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}

func main() {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", getEnv("PORT", "8080")),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  1 * time.Minute,
		Handler:      mux(),
	}
	log.Printf("Listening on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
