package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	Port = ":8081"
)

func serverDynamic(response http.ResponseWriter, request *http.Request) {
	responseText := "the time is now " + time.Now().String()
	//Fprintln allows you to direct output to any writer
	//println is only for console level
	fmt.Fprintln(response, responseText)
}

func serveStatic(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "static.html")
}

func main() {
	http.HandleFunc("/static", serveStatic)
	http.HandleFunc("/", serverDynamic)
	http.ListenAndServe(Port, nil)
}
