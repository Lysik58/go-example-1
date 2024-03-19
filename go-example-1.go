package main

import "fmt"

func makeCounter() func(i int) int {
	return func(i int) int {
		i++
		return i
	}
}
func main() {
	defer fmt.Println("End of the program")
	var counter = makeCounter()
	fmt.Println("Start of the program")

	for i := 0; i < 10; i++ {
		result := counter(0)
		if result > 5 {
			fmt.Println("Break")
			break
		}
		fmt.Println(result)
	}
}
