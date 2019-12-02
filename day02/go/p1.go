package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
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
	numbers[1] = 12
	numbers[2] = 2

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
			goto done
		default:
			log.Fatalf("Error while parsing code")
		}
	}

done:

	fmt.Printf("The final result for the first number = %d\n", numbers[0])
	return
}
