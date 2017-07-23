package main

import (
    "log"
    "net/http"
    "os"
    "fmt"

    "github.com/gorilla/mux"
)

type Pipe struct {
    ID    string `json:"id"`
    input string `json:"input"`
}

var pipes []Pipe

func indexHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
}

func inputHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf (w, "broooooooo")
}

func main() {

    PORT := os.Getenv("PORT")

    if PORT == "" {
        PORT = "3000"
    }

    router := mux.NewRouter()

    router.HandleFunc("/", indexHandler).Methods("GET")
    router.HandleFunc("/service/{id}", inputHandler).Methods("POST")
    log.Fatal(http.ListenAndServe(":" + PORT, router))

}
