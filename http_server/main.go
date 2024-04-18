package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "gabe.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := NewFileSystemPLayerStore(db)

	if err != nil {
		log.Fatalf("didn't expect an error but got one, %v", err)
	}

	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
