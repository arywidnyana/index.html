package main

import (
    "fmt"
)

func main() {
    var membership string
    var points [5]int

    fmt.Print("Masukkan status membership (Yes/No): ")
    fmt.Scan(&membership)
    
    fmt.Print("Masukkan lima poin yang diperoleh: ")
    for i := 0; i < 5; i++ {
        fmt.Scan(&points[i])
    }

    totalPoints := 0
    ganjil := true
    genap := true
    for i := 0; i < 5; i++ {
        totalPoints += points[i]
        if points[i] % 2 == 0 {
            ganjil = false
        } else {
            genap = false
        }
    }

    var cashback, diskon float64
    if ganjil {
        diskon = 1.7 * float64(points[2] + points[3] + points[4])
    } else if genap {
        cashback = 3.1 * float64(points[0] + points[1] + points[2])
    } else {
        cashback = 3.1 * float64(points[0] + points[1] + points[2])
        diskon = 1.7 * float64(points[2] + points[3] + points[4])
    }

    if membership == "Yes" {
        cashback += cashback * 0.15
        diskon += diskon * 0.15
    }

    if cashback > 35 {
        cashback = 35
    }
    if diskon > 50 {
        diskon = 50
    }

    fmt.Printf("Cashback: %.2f%%\n", cashback)
    fmt.Printf("Diskon: %.2f%%\n", diskon)
}
