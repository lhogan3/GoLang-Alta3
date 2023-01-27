package main

import (
    "fmt"
    "os"
)

func main() {
    //print out the number of arguments
    fmt.Println("Number of arguments:", len(os.Args[1:]))
}
