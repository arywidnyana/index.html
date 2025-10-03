package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Masukkan satu huruf: ")
	fmt.Scanln(&input)

	// Memastikan input satu karakter dan huruf
	if len(input) != 1 || !isLetter(input) {
		fmt.Println("bukan konsonan")
		return
	}

	// Cek apakah input adalah huruf konsonan
	if isConsonant(input) {
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan konsonan")
	}
}

// Fungsi untuk memeriksa apakah sebuah karakter adalah huruf
func isLetter(char string) bool {
	return (char >= "A" && char <= "Z") || (char >= "a" && char <= "z")
}

// Fungsi untuk memeriksa apakah sebuah huruf adalah konsonan
func isConsonant(char string) bool {
	// Huruf vokal
	vokal := "AEIOUaeiou"
	return !strings.Contains(vokal, char)
}
