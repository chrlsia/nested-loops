package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("i is", i)
		for j := 1; j <= 3; j++ {
			fmt.Println(" j:", j)
		}
		fmt.Println()
	}
}
