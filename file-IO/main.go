package main

import (
    "fmt"
    "os"
)

func main() {
    data, err := os.ReadFile("config.json")
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(data))
}