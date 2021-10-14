package main

import "fmt"

func playwithAsterik(number int) {
	for i := 1; i <= number; i++ {
		for j := number; j > i; j-- {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	playwithAsterik(5)
}
