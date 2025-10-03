package main

import (
    "fmt"
)

func main() {
    var suhu [5]float64
    fmt.Println("Masukkan lima temperatur:")
    for i := 0; i < 5; i++ {
        fmt.Scan(&suhu[i])
    }

    if isNaik(suhu) {
        fmt.Println("Stabil naik")
    } else if isTurun(suhu) {
        fmt.Println("Stabil turun")
    } else {
        fmt.Println("Tidak stabil")
    }
}

func isNaik(suhu [5]float64) bool {
    for i := 1; i < len(suhu); i++ {
        if suhu[i] <= suhu[i-1] {
            return false
        }
    }
    return true
}

func isTurun(suhu [5]float64) bool {
    for i := 1; i < len(suhu); i++ {
        if suhu[i] >= suhu[i-1] {
            return false
        }
    }
    return true
}
