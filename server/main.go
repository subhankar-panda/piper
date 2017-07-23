package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {

    PORT := os.Getenv("PORT")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World")
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":" + PORT, nil))

}
