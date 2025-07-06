package main

import (
	"fmt"
	"log"
	"net/http"
)

type database struct {
	db map[string]string
}

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db.db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	db := &database{
		db: map[string]string{"shoes": "50", "socks": "5"},
	}
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}
