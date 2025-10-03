package main

import (
    "fmt"
)

func main() {
    var a, b, c int
    fmt.Print("Masukkan tiga sisi segitiga: ")
    fmt.Scan(&a, &b, &c)

    if a == b && b == c {
        fmt.Println("Segitiga Sama Sisi")
    } else if a == b || b == c || a == c {
        fmt.Println("Segitiga Sama Kaki")
    } else {
        fmt.Println("Segitiga Sembarang")
    }
}
