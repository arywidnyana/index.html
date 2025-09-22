package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    correctUsername := "admin"
    correctPassword := "admin"
    attempts := 0

    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("Masukkan username: ")
        scanner.Scan()
        username := scanner.Text()

        fmt.Print("Masukkan password: ")
        scanner.Scan()
        password := scanner.Text()

        if username == correctUsername && password == correctPassword {
            fmt.Printf("Login berhasil setelah %d kali gagal login.\n", attempts)
            break
        } else {
            fmt.Println("Username atau password salah. Coba lagi.")
            attempts++
        }
    }
}
