package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func read_data() []int {
	intcode, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatalf("Error while getting line: %v", err)
	}

	var numbers []int
	for _, val := range strings.Split(intcode, ",") {
		if val == "" {
			continue
		}
		var code int
		_, err := fmt.Sscanf(val, "%d", &code)
		if err != nil {
			log.Fatalf("Error while getting number: %v", err)
		}
		numbers = append(numbers, code)
	}

	return numbers
}

func run_instructions(numbers []int) int {
	index := 0
	for {
		switch numbers[index] {
		case 1:
			numbers[numbers[index+3]] =
				numbers[numbers[index+1]] + numbers[numbers[index+2]]
			index += 4
		case 2:
			numbers[numbers[index+3]] =
				numbers[numbers[index+1]] * numbers[numbers[index+2]]
			index += 4
		case 99:
			return numbers[0]
		default:
			log.Fatalf("Error while parsing code")
		}
	}
}

func main() {
	orig_numbers := read_data()
	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			numbers := make([]int, len(orig_numbers))
			copy(numbers, orig_numbers)
			numbers[1] = noun
			numbers[2] = verb

			result := run_instructions(numbers)
			if result == 19690720 {
				fmt.Println("Result = ", (noun*100)+verb)
			}
		}
	}
}
