package main

import (
	"fmt"
)

func get_digits(num int) []int {
	digits := []int{}

	for num != 0 {
		digit := num % 10
		num /= 10
		digits = append(digits, digit)
	}

	return digits
}

func check_decreasing(digits []int) bool {
	for i := 0; i < 5; i++ {
		if digits[i] < digits[i+1] {
			return false
		}
	}

	return true
}

func check_exact_doubles(digits []int) bool {
	encountered := map[int]int{}
	for i := 0; i < 5; i++ {
		if digits[i] == digits[i+1] {
			encountered[digits[i]]++
		}
	}

	for _, v := range encountered {
		if v == 1 {
			return true
		}
	}

	return false
}

func check_number(num int) bool {
	digits := get_digits(num)
	return check_decreasing(digits) && check_exact_doubles(digits)
}

func main() {
	check := []int{341233, 555799, 111122, 111111, 112233, 123444, 123789}
	for _, i := range check {
		fmt.Println(i, check_number(i))
	}
	count := 0
	for i := 284639; i < 748759; i++ {
		if check_number(i) {
			count++
		}
	}

	fmt.Println("Count:", count)
}
