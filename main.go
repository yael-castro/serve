package main

import (
	"flag"
	"log"
	"net/http"
)

const (
	defaultPort = "8080"
	defaultDir  = "./"
)

func main() {
	port, dir := "", ""

	flag.StringVar(&port, "port", defaultPort, "bad port")
	flag.StringVar(&dir, "dir", defaultDir, "bad dir")

	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(dir)))

	log.Printf(`http server is running on "http://localhost:%s" 🐣`, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
