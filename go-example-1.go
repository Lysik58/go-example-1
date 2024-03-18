package main

import "fmt"

func makeCounter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	defer fmt.Println("End of the program")
	var counter = makeCounter()
	var result int
	fmt.Println("Start of the program")

	for i := 0; i < 10; i++ {
		result = counter()
		if result > 5 {
			fmt.Println("Break")
			break
		}
		fmt.Println(result)
	}
}
