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
		switch opcode {
		case 1:
			params := get_parameter(numbers, index, 3)
			numbers[numbers[index+3]] = params[0] + params[1]
			index += 4
		case 2:
			params := get_parameter(numbers, index, 3)
			numbers[numbers[index+3]] = params[0] * params[1]
			index += 4
		case 3:
			numbers[numbers[index+1]] = input
			index += 2
		case 4:
			output = numbers[numbers[index+1]]
			index += 2
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
	fmt.Println(run_instructions(intcode, 1))
}
