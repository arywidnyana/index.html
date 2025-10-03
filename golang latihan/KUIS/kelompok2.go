package main

import (
    "fmt"
)

func main() {
    var x int
    fmt.Print("Masukkan bilangan bulat positif: ")
    fmt.Scan(&x)

    if x <= 1 {
        fmt.Printf("%d bukan bilangan prima\n", x)
        return
    }

    isPrima := true
    for i := 2; i*i <= x; i++ {
        if x % i == 0 {
            isPrima = false
            break
        }
    }

    if isPrima {
        fmt.Printf("%d adalah bilangan prima\n", x)
    } else {
        fmt.Printf("%d bukan bilangan prima\n", x)
    }
}