package main

import (
	"log"
	"net/http"

	"github.com/Olamilekan-12/shortly/internal/server"
)

func main() {
	addr := ":8080"

	log.Printf("listening on port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.New()))
}
