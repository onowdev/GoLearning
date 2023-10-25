package main

import "fmt"

func main() {

	fmt.Println("Hasil Ujian Anda")

	testScore := 90
	testScoreGrade1 := 50
	testScoreGrade2 := 60
	testScoreGrade3 := 70
	testScoreGrade4 := 80

	if testScore == testScoreGrade1 {
		fmt.Println("Good Job")
	} else if testScore == testScoreGrade2 {
		fmt.Println("tingkatkan Terus Nilai nya")
	} else if testScore == testScoreGrade3 {
		fmt.Println("Ayo Bisa")
	} else if testScore > testScoreGrade4 {
		fmt.Println("Sempurna", testScore, "Nilai Anda")
	} else {
		fmt.Println("You Failed")
	}

}
