package main

import "fmt"

func main() {
	var studentScore int = 80
	var nilai string

	if studentScore >= 80 && studentScore <= 100 {
		nilai = "Nilai A"
	} else if studentScore >= 79 && studentScore <= 65 {
		nilai = "Nilai B+"
	} else if studentScore >= 64 && studentScore <= 50 {
		nilai = "Nilai B"
	} else if studentScore >= 49 && studentScore <= 35 {
		nilai = "Nilai C"
	} else if studentScore >= 34 && studentScore <= 0 {
		nilai = "Nilai D"
	}
	fmt.Println("Dapat Nilai :", nilai)
}
