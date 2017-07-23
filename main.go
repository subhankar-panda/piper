package main
import (

    "fmt"
    "os"
    "bufio"
    "strings"
    "io/ioutil"
    "math/rand"
    "time"

    flag "github.com/ogier/pflag"
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

        os.Exit(1)
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
    url := adjectiveArr[n]

    n = rand.Int() % len(nounArr)
    url += nounArr[n]

    return url
}

func init() {
    flag.StringVarP(&input, "input", "i", "", "Your input string")
    flag.StringVarP(&filepath, "file", "f", "", "Your source file")
}
