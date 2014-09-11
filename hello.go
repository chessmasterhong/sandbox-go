package main

import (
    "fmt"
    "time"
    "math/rand"
)

var t time.Time = time.Now()
var s int64 = t.Unix()

func main() {
    fmt.Println("Hello World!")

    fmt.Println("The time is", t)

    rand.Seed(s)
    var a, b, c int = sum(rand.Intn(5) + 1, rand.Intn(5) + 1)
    fmt.Println("I will roll two dice ... their values are", a, "and", b, "... their sum is", c)
}

func sum(x int, y int) (int, int, int) {
    return x, y, x + y
}
