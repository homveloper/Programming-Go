package main

import (
	"fmt"
	"log"
	"net/http"
)

type CFileServer struct {
}

func FormHandler(Writer http.ResponseWriter, Request *http.Request) {
	if err := Request.ParseForm(); err != nil {
		fmt.Fprintf(Writer, "ParseForm() Error %v", err)
	}

	fmt.Fprintf(Writer, "POST request successful")

	name := Request.FormValue("name")
	address := Request.FormValue("address")

	fmt.Fprintf(Writer, "name = %s\n", name)
	fmt.Fprintf(Writer, "address = %s\n", address)
}

func HelloHandler(Writer http.ResponseWriter, Request *http.Request) {
	if Request.URL.Path != "/hello" {
		http.Error(Writer, "404 Not Found", http.StatusNotFound)
		return
	}

	if Request.Method != "GET" {
		http.Error(Writer, "Method Is Not Supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(Writer, "Hello")
}

func main() {
	FileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", FileServer)
	http.HandleFunc("/form", FormHandler)
	http.HandleFunc("/hello ", HelloHandler)

	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
