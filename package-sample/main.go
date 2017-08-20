package main
import (
    "fmt"
    "os"
    "os/exec"
    "bufio"
    "./classes"
)

func main() {
    fmt.Print("Hello world from test!\n")
    stdin := bufio.NewScanner(os.Stdin)
    for stdin.Scan() {
        test := stdin.Text()
        fmt.Println(test)
    }
}
