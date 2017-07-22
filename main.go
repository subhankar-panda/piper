package main
import (

    "fmt"
    "os"
    "bufio"

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

        fmt.Println(input)
        os.Exit(1)
    }
}

func init() {
    flag.StringVarP(&input, "input", "i", "", "Your input string")
    flag.StringVarP(&filepath, "file", "f", "", "Your source file")
}
