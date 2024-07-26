package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello_world", handleHelloWorld)
	http.HandleFunc("/health", handleHealth)

	addr := "localhost:8000"
	log.Printf("Listening on %s ...", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleHelloWorld(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	writeResponse(writer, "Hello World")
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	writeResponse(writer, "Ok!")
}

func writeResponse(writer http.ResponseWriter, responseString string) {
	response := []byte(responseString) // slice of byte i.e.  [72 101 108 108 111 32 87 111 114 108 100 33]
	fmt.Println(response)
	_, err := writer.Write(response)

	if err != nil {
		fmt.Println(err)
	}
}