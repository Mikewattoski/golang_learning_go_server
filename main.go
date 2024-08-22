package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./Static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Wrong method", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Hello!!!")
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Post request success!!")
	name := r.FormValue("name")
	pnumber := r.FormValue("pnumber")
	email := r.FormValue("Email")
	website := r.FormValue("website")
	fmt.Fprintf(w, "Name:%v,Phone number:%v,Email:%v,Website:%v", name, pnumber, email, website)
}
