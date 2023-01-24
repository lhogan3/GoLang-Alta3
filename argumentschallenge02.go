package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]

    for i, a := range args {
        fmt.Println(a, "was passed in position", i)
    }
}
