package main

import (
    "fmt"
    "math"
    //"math/rand"
    "strconv"
    //"time"
)

const PI = math.Pi
var pi float64 = PI

func main() {
//    fmt.Println("Hello World!")
//
//    var t time.Time = time.Now()
//    fmt.Println("The time is", t)
//
//    rand.Seed(t.Unix())
//    var a, b, c int = sum(rand.Intn(5) + 1, rand.Intn(5) + 1)
//    fmt.Println("I will roll two dice ... their values are", a, "and", b, "... their sum is", c)
//
//    var pi float64 = float64(int(pi * 100)) / 100
//    fmt.Println("Happy", pi, "Day!")
//
//    for i := 10; i > 0; i-- {
//        fmt.Println(i, "bottles of beer on the wall,", i, "bottles of beer. Take one down pass it around,", i - 1, "bottles of beer on the wall.")
//    }
    
    for i := 1; i <= 100; i++ {
        fmt.Printf(strconv.Itoa(i) + " ")

        if i % 3 == 0 {
            fmt.Printf("Fizz")
        }
        if i % 5 == 0 {
            fmt.Printf("Buzz")
        }

        fmt.Printf("\n")
    }
}

func sum(x int, y int) (int, int, int) {
    return x, y, x + y
}
