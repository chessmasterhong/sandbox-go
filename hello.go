package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    fmt.Println("Hello World!")

    fmt.Println("The time is", time.Now())

    rand.Seed(time.Now().Unix())
    fmt.Println("I will roll two dices... their values are", rand.Intn(6), "and", rand.Intn(6))
}
