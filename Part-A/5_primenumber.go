package main

import "fmt"

func primeNumber(number int) bool {
	var faktor int
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			faktor++
		}
	}
	if faktor == 2 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(primeNumber(11))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
