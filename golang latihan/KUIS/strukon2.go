package main

import "fmt"

func main() {
	// Input dua bilangan
	var n, m int
	fmt.Print("Masukkan bilangan pertama (N): ")
	fmt.Scan(&n)
	fmt.Print("Masukkan bilangan kedua (M): ")
	fmt.Scan(&m)

	var fpb int
	for i := 1; i <= n && i <= m; i++ {
		if n%i == 0 && m%i == 0 {
			fpb = i
		}
	}
	
	// Cari KPK
	kpk := (n * m) / fpb
	fmt.Println("Kelipatan Persekutuan Terkecil (KPK):", kpk)
}