package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
)

type Pipe struct {
    Input string  `json:"input"`
    URL   string  `json:"url"`
}

func main() {

    PORT := os.Getenv("PORT")

    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/", indexHandler)

    log.Fatal(http.ListenAndServe(":" + PORT, nil))

}
