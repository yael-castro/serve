package main

import (
	"flag"
	"log"
	"net/http"
)

const (
	defaultPort = "80"
	defaultDir  = "./public"
)

func main() {
	port, dir := "", ""

	flag.StringVar(&port, "port", defaultPort, "bad port")
	flag.StringVar(&dir, "dir", defaultDir, "bad dir")

	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(dir)))

	log.Printf(`Files in directory "%s" are served at "http://localhost:%s"`, dir, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
