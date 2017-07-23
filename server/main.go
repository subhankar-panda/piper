package main

import (
    "log"
    "net/http"
    "os"
    "fmt"
    "strings"
    "io/ioutil"
    "encoding/json"

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
    body, err := ioutil.ReadAll(r.Body)

    var pipe Pipe

    if err != nil {
        panic(err)
        print(body)
    }

    jsonString := string(body)

    json.Unmarshal([]byte(jsonString), &pipe)
    fmt.Fprintf(w, pipe.input)
}


func formatRequest(r *http.Request) string {

 var request []string

 url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
 request = append(request, url)

 request = append(request, fmt.Sprintf("Host: %v", r.Host))

 for name, headers := range r.Header {
   name = strings.ToLower(name)
   for _, h := range headers {
     request = append(request, fmt.Sprintf("%v: %v", name, h))
   }
 }


 if r.Method == "POST" {
    r.ParseForm()
    request = append(request, "\n")
    request = append(request, r.Form.Encode())
 }

  return strings.Join(request, "\n")
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
