package main

import "fmt"

func makeCounter(a int) func() int {

	return func() int {
		a++
		return a
	}
}
func main() {
	defer fmt.Println("End of the program")
	var counter = makeCounter(0)
	fmt.Println("Start of the program")

	for i := 0; i < 10; i++ {
		result := counter()
		fmt.Println(result)
		if result > 5 {
			fmt.Println("Break")
			break
		}
	}
}
