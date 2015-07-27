package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Local()
	_, err := fmt.Fprintf(w, "Current time : ")
	_, err1 := fmt.Fprintf(w, t.Format("2006-01-02 15:04:05"))
	if err != nil || err1 != nil {
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
	for {
		fmt.Println("Please input port")
		_, err2 := fmt.Scanf("%s", &s)
		if err2 == nil {
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
			if err3 := http.ListenAndServeTLS(s, "cert.pem", "key.pem", nil); err3 != nil {
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