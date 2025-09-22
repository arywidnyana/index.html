package main

import "fmt"

func main() {
	var n, m int
	fmt.Print("Masukkan bilangan pertama (N): ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan kedua (M): ")
	fmt.Scan(&m)

	var faktor []int
	for i := 1; i <= n && i <= m; i++ {
		if n%i == 0 && m%i == 0 {
			faktor = append(faktor, i)
		}
	}

	fmt.Println("Faktor Persekutuan:", faktor)
}