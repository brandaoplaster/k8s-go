package main

import "net/http"

func main() {
	println("Hello, World!")

	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Hello(write http.ResponseWriter, request *http.Request) {
	write.Write([]byte("Hello, World!"))
}
