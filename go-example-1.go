package main

import "fmt"

func main() {
	defer fmt.Println("End of the program")
	var result int
	fmt.Println("Start of the program")
	//var counter = makeCounter(0)

	for i := 0; i < 10; i++ {
		result = counter
		if result > 5 {
			fmt.Println("Break")
			break
		}
		fmt.Println(result)
	}
}
