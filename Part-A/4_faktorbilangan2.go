package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Input :")
	fmt.Scanln(&bilangan)

	for i := bilangan; i > 0; i-- {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}
}
