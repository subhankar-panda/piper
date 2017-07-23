package main

import (
  //  "encoding/json"
    "fmt"
    "net/http"

//    "github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello World!")
}
