package main

import (
    "fmt"
)

func main() {
    var x int
    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&x)

    fmt.Printf("Faktor dari %d adalah: ", x)
    for i := 1; i <= x; i++ {
        if x % i == 0 {
            fmt.Printf("%d ", i)
        }
    }
    fmt.Println()
}