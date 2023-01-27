package main

import "fmt"

func main() {
    pl := map[string]string{"Python": ".py", "Golang": ".go", "Java": ".java", "Ansible": ".yml", "C++": ".cpp"}
    fmt.Println("Full Programming Languages Map:", pl)

    delete(pl, "C++")
    pl["Julia"] = ".jl"
    fmt.Println("Modified Programming Languages Map:", pl)
}
