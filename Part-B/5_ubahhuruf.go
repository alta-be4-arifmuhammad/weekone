package main

import (
	"fmt"
)

func ubahHuruf(sentence string) string {
	newSentence := []byte(sentence)
	for i := 0; i < len(sentence); i++ {
		if sentence[i] >= 'A' && sentence[i] <= 'P' {
			newSentence[i] = sentence[i] + 10
		} else if sentence[i] >= 'Q' && sentence[i] <= 'Z' {
			newSentence[i] = sentence[i] - 16
		}
	}
	return string(newSentence)
}

func main() {
	fmt.Println(ubahHuruf("SEPULSA OKE"))
}
