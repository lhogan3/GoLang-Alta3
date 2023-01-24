package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "os"
)

func main() {
    var numLines int
    flag.IntVar(&numLines, "n", 5, "number of lines to be displayed from the top of the file")
    flag.Parse()

    var inFile io.Reader
    if filename := flag.Arg(0); filename != "" {
        file, err := os.Open(filename)

        if err != nil {
            fmt.Println("file not able to be opened with the following error:", err)
            os.Exit(1)
        }
        defer file.Close()

        inFile = file
    } else {
        fmt.Println("Filename to read from not provided...")
        inFile = os.Stdin
    }

    buf := bufio.NewScanner(inFile)

    for i := 0; i < numLines; i++ {
        if !buf.Scan() {
            break
        }
        fmt.Println(buf.Text())
    }

    if err := buf.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "file not able to be read: ", err)
    }
}
