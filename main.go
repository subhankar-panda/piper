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
    "github.com/fatih/color"
)

const (
   API_URI = "https://agile-escarpment-29641.herokuapp.com/"
//   API_URI = "http://localhost:3000/"
    OUTPUT_STR = "Yay! Your output is now live at \n \t %s%s"
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
        currtime := time.Now().String()
        sending := map[string]string{"id" : ext, "input" : input, "time" : currtime}
        JSON, err := json.Marshal(sending)

        if err != nil {
            panic(err)
        }

        req, err := http.NewRequest("POST", API_URI + "service/" + ext , bytes.NewBuffer(JSON))

        if err != nil {
            panic(err)
        }

        client := &http.Client{}
        resp, err := client.Do(req)

        if err != nil {
            panic(err)
        }

        defer resp.Body.Close()

        _, err = ioutil.ReadAll(resp.Body)

        if err != nil {
            panic(err)
        }

        red := color.New(color.FgRed).SprintFunc()

        fmt.Printf(OUTPUT_STR, API_URI ,red(ext))
        fmt.Println("")
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
