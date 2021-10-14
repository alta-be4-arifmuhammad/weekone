package main

import (
	"fmt"
)

func cetakTablePerkalian(number int) {
	for rows := 1; rows <= number; rows++ {
		for column := 1; column <= number; column++ {
			total := rows * column
			fmt.Printf("%d\t", total)
		}
		fmt.Println("\n")
	}
}

func main() {
	cetakTablePerkalian(9)
}
