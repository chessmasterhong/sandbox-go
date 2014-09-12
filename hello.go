package main

import (
    "fmt"
    "math"
    "math/rand"
    "strconv"
    "time"
)

const PI = math.Pi
var pi float64 = PI

func main() {
    fmt.Println("Hello World!")

    var t time.Time = time.Now()
    fmt.Println("The time is", t)

    rand.Seed(t.Unix())
    var a, b, c int = sum(rand.Intn(5) + 1, rand.Intn(5) + 1)
    fmt.Println("I will roll two dice ... their values are", a, "and", b, "... their sum is", c)

    var pi float64 = float64(int(pi * 100)) / 100
    fmt.Println("Happy", pi, "Day!")

    //BeerBottles(99)

    //FizzBuzz(100)
}

func sum(x int, y int) (int, int, int) {
    return x, y, x + y
}

func BeerBottles(b int) {
    if b == 0 {
        return
    }

    fmt.Println(b, "bottles of beer on the wall,", b, "bottles of beer. Take one down, pass it around,", b - 1, "bottles of beer on the wall.")

    BeerBottles(b - 1)
}

func FizzBuzz(n int) {
    if n == 0 {
        return
    }

    FizzBuzz(n - 1)

    fmt.Printf(strconv.Itoa(n) + " ")

    if n % 3 == 0 {
        fmt.Printf("Fizz")
    }
    if n % 5 == 0 {
        fmt.Printf("Buzz")
    }

    fmt.Printf("\n")
}
