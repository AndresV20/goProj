package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err!=nil {
		fmt.Printf(w, "ParseForm() err: %v")
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "method is not supported", http.StatusNotFound)
		return 
	}
	fmt.Printf(w, "hello!")
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", hellohandler)


	fmt.Printf("Starting server at port 8080")
	if err:= http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatal(err)
	}
}