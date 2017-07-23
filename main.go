package main
import (

    "fmt"
    "os"
    "bufio"
    "strings"
    "io/ioutil"
    "math/rand"
    "time"
    "encoding/json"
    "net/http"
    "bytes"

    flag "github.com/ogier/pflag"
)

const (
    API_URI = "https://agile-escarpment-29641.herokuapp.com/"
)

var (
    input string
    filepath string
)

func main() {

    flag.Parse()

    info, _ := os.Stdin.Stat()

    if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
        fmt.Println("The command is intended to work with pipes.")
        fmt.Println("Usage:")
        fmt.Println("  cat yourfile.txt | piper <your_code_snippet>")
        os.Exit(0)

    } else if info.Size() > 0 {
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            input += scanner.Text() + "\n"
        }

        ext := createURL()
        sending := map[string]string{"id" : ext, "input" : input}
        JSON, _ := json.Marshal(sending)
        req, _ := http.NewRequest("POST", API_URI + "service/" + ext , bytes.NewBuffer(JSON))

        client := &http.Client{}
        resp, err := client.Do(req)

        if err != nil {
            panic(err)
        }

        defer resp.Body.Close()

        body, _ := ioutil.ReadAll(resp.Body)
        fmt.Println(string(body))
    }
}

func createURL() (words string) {

    rand.Seed(time.Now().Unix())

    adjectives, _ := os.Open("./words/adjectives.txt")
    nouns, _ := os.Open("./words/nouns.txt")

    bytesAdj, _ := ioutil.ReadAll(adjectives)
    bytesNoun, _  := ioutil.ReadAll(nouns)

    adjectiveArr := strings.Split(string(bytesAdj), "\n")
    nounArr := strings.Split(string(bytesNoun), "\n")

    n := rand.Int() % len(adjectiveArr)
    url := strings.Title(adjectiveArr[n])

    n = rand.Int() % len(nounArr)
    url += strings.Title(nounArr[n])

    return url
}

func init() {
    flag.StringVarP(&input, "input", "i", "", "Your input string")
    flag.StringVarP(&filepath, "file", "f", "", "Your source file")
}
