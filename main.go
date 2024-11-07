package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Log the current working directory to confirm it's correct
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Serving files from:", dir)

	// Serve files from the 'my files' directory
	fs := http.FileServer(http.Dir("./my files"))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(":80", nil))
}
