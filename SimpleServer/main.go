package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("received request: ", r.RemoteAddr)
		now := time.Now().Format("2006-01-02 15:04:05.000")
		responseText := fmt.Sprintf("Hello Docker!\nNow: %s\n", now)
		fmt.Fprintf(rw, "%s", responseText)
	})

	log.Println("start server")
	server := &http.Server{Addr: ":8080"}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
