package main

import (
    "strings"
    "io/ioutil"
    "encoding/json"
    "net/http"
    "fmt"
    "os"
    "log"
    "gopkg.in/mgo.v2"

    "github.com/gorilla/mux"
)

type Pipe struct {
    ID    string `json:"id"`
    Input string `json:"input"`
}

var pipes []Pipe

func indexHandler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
}

func inputHandler(w http.ResponseWriter, req *http.Request) {
    body, err := ioutil.ReadAll(req.Body)

    var pipe Pipe

    if err != nil {
        panic(err)
    }

    err = req.Body.Close()

    if err != nil {
        panic(err)
    }

    err = json.Unmarshal (body, &pipe)

    if err != nil {
        panic(err)
    }

    creds, _ := ioutil.ReadFile("creds.txt")

    credsString := strings.TrimSpace(string(creds))

    if credsString == "" {
        credsString = os.Getenv("PASS")
    }

    if err != nil {
        panic(err)
    }

    uri := "mongodb://subhankarpanda:" + credsString + "@ds047592.mlab.com:47592/piper"

    if uri == "" {
        fmt.Println("no connection string provided")
        os.Exit(1)
    }

    sess, err := mgo.Dial(uri)

    if err != nil {
        fmt.Printf("Can't connect to mongo, go error %v\n", err)
        os.Exit(1)
    }

    defer sess.Close()

    sess.SetSafe(&mgo.Safe{})

    c := sess.DB("piper").C("pipes")

    err = c.Insert(&Pipe{ID: pipe.ID, Input: pipe.Input})

    if err != nil {
        panic(err)
    }

    fmt.Fprintln(w, pipe.ID)
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

func getValueFunc(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    fmt.Fprintln(w, params)
}

func main() {

    PORT := os.Getenv("PORT")

    if PORT == "" {
        PORT = "3000"
    }

    router := mux.NewRouter()

    router.HandleFunc("/", indexHandler).Methods("GET")
    router.HandleFunc("/service/{id}", inputHandler).Methods("POST")
    router.HandleFunc("/{id}", getValueFunc).Methods("GET")
    log.Fatal(http.ListenAndServe(":" + PORT, router))

}
