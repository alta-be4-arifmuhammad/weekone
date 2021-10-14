package main

import "fmt"

func DrawXYZ(high int) string {
	var number int
	str := ""
	for i := 1; i <= high; i++ {
		for j := 1; j <= high; j++ {
			number++
			if number%3 == 0 {
				str = str + "X "
			} else if number%2 == 0 {
				str = str + "Z "
			} else {
				str = str + "Y "
			}
		}
		str = str + "\n"
	}
	return str
}

func main() {
	fmt.Println(DrawXYZ(3))
	fmt.Println(DrawXYZ(5))
}
