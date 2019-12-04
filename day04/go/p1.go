package main

import (
	"fmt"
)

func check_number(num int) bool {
	has_consecutive := false
	last_digit := num % 10
	num /= 10
	for num != 0 {
		digit := num % 10
		num /= 10

		if digit > last_digit {
			return false
		}

		if digit == last_digit {
			has_consecutive = true
		}

		last_digit = digit
	}

	return has_consecutive
}

func main() {
	fmt.Println(111111, check_number(111111))
	fmt.Println(223450, check_number(223450))
	count := 0
	for i := 284639; i <= 748759; i++ {
		if check_number(i) {
			count++
		}
	}

	fmt.Println("Count:", count)
}
