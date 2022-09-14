package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zekapp/multi-language-bazel-monorepo/projects/go_hellow_world"
)


func YourHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request")
	w.Write([]byte(go_hello_world.HelloWorld()))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/" YourHandler)

	log.Println("Going to listen on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}