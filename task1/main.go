package main

import (
	"fmt"
	"math"
)

func getDigitsCount(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}
	return count
}

func main() {
	fmt.Print("Enter your number: ")
	var number int
	fmt.Scanln(&number)
	fmt.Println("Your number's square is:", number*number)
	if number*number%int(math.Pow(10, float64(getDigitsCount(number)))) == number {
		fmt.Println("Yes,", number, "is automorphic")
	} else {
		fmt.Println("No,", number, "is not automorphic")
	}
}
