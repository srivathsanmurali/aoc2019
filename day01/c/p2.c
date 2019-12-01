#include <stdio.h>
#include <math.h>

int calculate_fuel_cost(int mass) {
  int fuel = floor(mass/3) - 2;
  if(fuel <= 0) {
    return 0;
  } else {
    return fuel + calculate_fuel_cost(fuel);
  }
}

int main() {
  int totalFuel = 0;

  while(1) {
    int moduleMass;
    if(scanf("%d", &moduleMass) != EOF) {
      totalFuel += calculate_fuel_cost(moduleMass);
    } else {
      break;
    }
  }

  printf("Total fuel required: %d\n", totalFuel);
}
