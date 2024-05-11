package main

import (
	"fmt"
	"net/http"
	"time"
)

var startedAt = time.Now()

func main() {
	println("Hello, World!")

	http.HandleFunc("/", Hello)
	http.HandleFunc("/healthz", Healthz)
	http.ListenAndServe(":80", nil)
}

func Hello(write http.ResponseWriter, request *http.Request) {
	write.Write([]byte("Hello, World!"))
}

func Healthz(w http.ResponseWriter, r *http.Request) {

	duration := time.Since(startedAt)

	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}

}
