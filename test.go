package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	const layout = "3:04pm"
	t := time.Now()
	if _, err1 := fmt.Fprintln(w, t.UTC().Format(layout)); err1 != nil {
		log.Fatal("failed to write", err1)
	}
}

func handlerCMDArgs() {
	port := flag.String("port", ":9111", "port in server")
	flag.Parse()
	if err3 := http.ListenAndServeTLS(*port, "cert.pem", "key.pem", nil); err3 != nil {
		log.Fatal("failed to start server", err3)
	}
}

func main() {
	http.HandleFunc("/time", handler)
	handlerCMDArgs()
}