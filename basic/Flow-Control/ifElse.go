package main

import "fmt"

func main() {

	accountBalance := 900
	accountCredit := 8000

	// Mengunakan if

	if accountBalance+accountCredit > 500 {
		fmt.Println("Uang anda lebih dari 500 $")
	}

	// Menggunakan if else
	if accountBalance+accountCredit < 0 {
		fmt.Println("Saldo anda kurang, Ayo Tambahkan Saldo Anda")
	} else {
		fmt.Println("Anda Punya Uang")
	}

}
