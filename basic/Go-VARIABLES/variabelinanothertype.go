package main

import "fmt"

func main() {

	var (
		firstName = "Chris"
		age       = 25
	)

	fmt.Println("Nama Lengkap = ", firstName)
	fmt.Println("Usia = ", age)

	var (
		customerName   = "Andra"
		accountBalance = 9000000
	)

	fmt.Println("Customer %s has %d$ on their Bank account", customerName, accountBalance)

}
