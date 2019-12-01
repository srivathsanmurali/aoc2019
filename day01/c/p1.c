#include <stdio.h>
#include <math.h>

int main() {
  int totalFuel = 0;

  while(1) {
    int moduleMass;
    if(scanf("%d", &moduleMass) != EOF) {
      totalFuel += floor(moduleMass/3) - 2;
    } else {
      break;
    }
  }

  printf("Total fuel required: %d\n", totalFuel);
}
