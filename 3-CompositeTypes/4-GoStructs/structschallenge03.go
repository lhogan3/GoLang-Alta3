package main

import "fmt"

type Astro struct {
    name     string
    age      int
    mission  string
    isneeded bool
}

type nasaMission struct {
    people []Astro
    number int
    message string
}

func main() {
    astro1 := Astro{"Liam", 24, "Go to the Moon", true}
    astro2 := Astro{"Buzz", 30, "Get home safe", false}
    astro3 := Astro{"LeBron", 38, "Go to the farthest", false}

    fmt.Println(astro1)
    fmt.Println(astro2)
    fmt.Println(astro3)

    p := []Astro{astro1, astro2, astro3}

    s := nasaMission{p, 3, "succes"}

    fmt.Println(s)
    fmt.Printf("%+v", s)
}
