package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	totalFuel := 0

	for scanner.Scan() {
		var moduleMass int

		_, err := fmt.Sscanf(scanner.Text(), "%d", &moduleMass)
		if err != nil {
			log.Fatalf("Error while getting number: %v", err)
		}

		totalFuel += int(moduleMass/3) - 2
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Can't read lines %v", err)
	}

	fmt.Printf("Total fuel required: %d\n", totalFuel)

	return
}
