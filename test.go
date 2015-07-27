package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Local()
	_, err1 := fmt.Fprintf(w, t.Format("2006-01-02 15:04:05"))
	if err1 != nil {
		log.Fatal("failed to write", err, err1)
	}
	defer func() {
		if x := recover(); x != nil {
			log.Printf("%s", x, "caught panic")
		}
	}()
}

func input() {
	var s string
	a := " "
	for {
		s = " "
		if len(os.Args) > 1 {
			s = os.Args[1]
		}
		if s != a {
			k, y := regexp.MatchString("[0-9]", s)
			if k == true && y == nil {
				s = ":" + s
			}
			if err3 := http.ListenAndServeTLS(s, "cert.pem", "key.pem", nil); err3 != nil {
				log.Println("failed to start server", err3)
				continue
			} else {
				break
			}
		} else {
			if err3 := http.ListenAndServeTLS(":9876", "cert.pem", "key.pem", nil); err3 != nil {
				log.Println("failed to start server", err3)
				continue
			}
		}
	}
}

func main() {
	http.HandleFunc("/gettime", handler)
	input()
}