package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	t := time.Now()
	_, err1 := fmt.Fprintf(w, t.UTC().Format(layout))
	if err1 != nil {
		log.Fatal("failed to write", err1)
	}
	defer func() {
		if x := recover(); x != nil {
			log.Printf("%s", x, "caught panic")
		}
	}()
}

func input() {
	var s string
	port := flag.String("port", "9111", "port in server")
	flag.Parse()
	s = *port
	k, y := regexp.MatchString("[0-9]", s)
	if k == true && y == nil {
		s = ":" + s
		if err3 := http.ListenAndServeTLS(s, "cert.pem", "key.pem", nil); err3 != nil {
			log.Println("failed to start server", err3)
		} else {

		}
	} else {
		log.Println("unexpected input", y)
	}
}

func main() {
	http.HandleFunc("/time", handler)
	input()
}
