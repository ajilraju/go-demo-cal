package maine

import (
    "fmt"
    "log"
    "os"
)

var VERSION string = "1.0.3"

var (
    BOLD   string = "\033[1m"
    YELLOW string = "\033[93m"
    BLUE   string = "\033[94m"
    END    string = "\033[0m"
)

func add(f int, s int) int {
    return f + s
}

func multi(f int, s int) int {
    return f * s
}

func sub(f int, s int) int {
    return f - s
}

func div(f int, s int) int {
    return f / s
}

func main() {

    var firstNum int
    var secondNum int

    if len(os.Args) == 2 {
        if os.Args[1] == "-v" {
            fmt.Printf("%sAncient simple Calculator: %s%s\n", BOLD, VERSION, END)
            return
        } else {
            return
        }
    }

    fmt.Printf("%sAncient simple Calculator%s\n\n", BOLD, END)

    fmt.Println("Enter the first digit: ")
    if _, err := fmt.Scan(&firstNum); err != nil {
        log.Print("Failed to scan first number", err)
        return
    }

    fmt.Println("Enter the second digit")
    if _, err := fmt.Scan(&secondNum); err != nil {
        log.Print("Failed to scan second number", err)
        return
    }

    fmt.Printf("%sSum     [ %d + %d = %d ]%s\n", BLUE, firstNum, secondNum, add(firstNum, secondNum), END)
    fmt.Printf("%sProduct [ %d * %d = %d ]%s\n", YELLOW, firstNum, secondNum, multi(firstNum, secondNum), END)
    fmt.Printf("%sDifference [ %d - %d = %d ]%s\n", BLUE, firstNum, secondNum, sub(firstNum, secondNum), END)
    fmt.Printf("%sQuotient [ %d / %d = %d ]%s\n", YELLOW, firstNum, secondNum, div(firstNum, secondNum), END)
}
