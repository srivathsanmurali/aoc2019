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

		totalFuel += calculate_fuel_cost(moduleMass)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Can't read lines %v", err)
	}

	fmt.Printf("Total fuel required: %d\n", totalFuel)

	return
}

func calculate_fuel_cost(mass int) int {
	fuel := int(mass/3) - 2
	if fuel <= 0 {
		return 0
	} else {
		return fuel + calculate_fuel_cost(fuel)
	}
}
