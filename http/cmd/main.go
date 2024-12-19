package main

import (
	http2 "learnGoWithTests/http"
	"learnGoWithTests/http/storage/memStorage"
	"log"
	"net/http"
)

func main() {
	store := memStorage.NewMemStorage()
	server := http2.PlayerServer{Store: store}

	log.Fatal(http.ListenAndServe(":5000", &server))
}
