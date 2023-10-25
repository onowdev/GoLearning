package main

import "fmt"

func main() {

	var exsist bool = true
	fmt.Println("Exsist? %t \n", exsist)

	printMessage := true

	if printMessage {
		fmt.Println("Pesan")
		fmt.Println("Ayo Belajar")
	}

	accountBallance := 800
	accountCredit := 9000

	if accountBallance+accountCredit > 2000 {
		fmt.Println("Anda Punya uang untuk di belanja kan")
	}
}
