package main

import (
	"PulseStream-MacAudio/lib"
	"log"
	"net/http"
)

func main() {
	http.Handle("/ffmpeg", new(lib.CountHandler))

	fs := http.FileServer(http.Dir("./pages"))
	http.Handle("/", fs)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
