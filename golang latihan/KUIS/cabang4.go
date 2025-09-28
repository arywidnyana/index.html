package main

import (
    "fmt"
    "math"
)

func main() {
    var input int
    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&input)
    output := int(math.Abs(float64(input)))
    fmt.Printf("Nilai mutlak: %d\n", output)
}
