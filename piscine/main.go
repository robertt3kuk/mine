package main

import "fmt"

func main() {
	fmt.Println(prime(12345436))
	fmt.Println(prmu(12345436))
	fmt.Println(1)
}
func prime(number int) int {
	if number == 1 {
		return 0
	}
	if number%2 != 0 && number%3 != 0 && number%5 != 0 && number%7 != 0 {
		return number
	} else if number == 2 {
		return 2
	} else if number == 3 {
		return 3
	} else {
		return prime(number - 1)
	}
}
func prmu(number int) int {
	for number != number/2 {
		if number%2 != 0 && number%3 != 0 && number%5 != 0 && number%7 != 0 {
			return number
		}
		number--
	}
	return 0
}
