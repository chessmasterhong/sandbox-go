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
    a, b, c := sum(rand.Intn(5) + 1, rand.Intn(5) + 1)
    fmt.Println("I will roll two dice ... their values are", a, "and", b, "... their sum is", c)
}

func sum(x int, y int) (int, int, int) {
    return x, y, x + y
}
