package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func read_intcode() []int {
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

/* get_parameter
 * instructions: The set of instructions
 * index: The index of the instruction
 */
func get_parameter(numbers []int, index int, max_params int) []int {
	parameter_instruction := numbers[index] / 100
	parameters := []int{}
	for i := 1; i <= max_params; i++ {
		code := parameter_instruction % 10
		parameter_instruction /= 10
		switch code {
		case 0:
			parameters = append(parameters, numbers[numbers[index+i]])
		case 1:
			parameters = append(parameters, numbers[index+i])
		}
	}

	return parameters
}

func run_instructions(numbers []int, input int) int {
	index := 0
	output := 0
	for {
		opcode := numbers[index] % 100
		fmt.Printf("index: %d, code: %d, opcode: %d\n", index, numbers[index], opcode)
		switch opcode {
		case 1:
			// add
			params := get_parameter(numbers, index, 3)
			numbers[numbers[index+3]] = params[0] + params[1]
			index += 4
		case 2:
			// multiply
			params := get_parameter(numbers, index, 3)
			numbers[numbers[index+3]] = params[0] * params[1]
			index += 4
		case 3:
			// input
			numbers[numbers[index+1]] = input
			index += 2
		case 4:
			// output
			output = numbers[numbers[index+1]]
			index += 2
		case 5:
			// jump-if-true
			params := get_parameter(numbers, index, 2)
			if params[0] != 0 {
				index = params[1]
			} else {
				index += 3
			}
		case 6:
			// jump-if-false
			params := get_parameter(numbers, index, 2)
			if params[0] == 0 {
				index = params[1]
			} else {
				index += 3
			}
		case 7:
			// less than
			params := get_parameter(numbers, index, 2)
			if params[0] < params[1] {
				numbers[numbers[index+3]] = 1
			} else {
				numbers[numbers[index+3]] = 0
			}
			index += 4
		case 8:
			// equals to
			params := get_parameter(numbers, index, 2)
			if params[0] == params[1] {
				numbers[numbers[index+3]] = 1
			} else {
				numbers[numbers[index+3]] = 0
			}
			index += 4
		case 99:
			return output
		default:
			log.Fatalf("Error while parsing code: %v", opcode)
		}
	}
}

func main() {
	// Read intcode from stdin
	intcode := read_intcode()
	fmt.Printf("Length = %d\n", len(intcode))

	// Run the Instructions
	fmt.Println(run_instructions(intcode, 5))
}
