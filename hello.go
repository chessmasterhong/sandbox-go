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
    fmt.Println("I will roll two dice... the sum of the dice are", sum(rand.Intn(6), rand.Intn(6)))
}

func sum(x int, y int) int {
    return x + y
}