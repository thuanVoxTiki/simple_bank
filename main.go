package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	fmt.Println(sum(5))
}

func sum(num int) int {
	sum := 0
	for i := 0; i <= num; i++ {
		sum += i
	}
	fmt.Printf("%T", num)
	fmt.Println()
	return sum
}
