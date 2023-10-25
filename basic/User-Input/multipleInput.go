package main

import "fmt"

func main() {

	var namaDepan, namaBelakang string
	fmt.Scan(&namaDepan, &namaBelakang)

	//formated String to print out the User Input

	str := fmt.Sprintf("Nama Depan %s Nama Belakang %s", namaDepan, namaBelakang) //("Nama Depan = %s", namaDepan, "Nama Belakang = ", namaBelakang)

	fmt.Println(str)

}
