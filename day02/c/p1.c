#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int* get_instructions() {
  int* memory = malloc(sizeof(int) * 128);
  int address = 0;
  while(1) {
    int code = 0;
    if(scanf("%d,", &code) != EOF) {
      memory[address++] = code;
    } else {
      break;
    }
  }

  return memory;
}

int run(int* memory) {
  int* ip = memory;
  while(1) {
    switch(*ip) {
      case 1:
        memory[*(ip+3)] = memory[*(ip+1)] + memory[*(ip+2)];
        break;
      case 2:
        memory[*(ip+3)] = memory[*(ip+1)] * memory[*(ip+2)];
        break;
      case 99:
        return memory[0];
      default:
        return -1;
    }
    ip += 4;
  }
}

int main() {
  int* memory = get_instructions();
  memory[1] = 12;
  memory[2] = 2;
  int result = run(memory);
  printf("Result: %d\n", result);
  free(memory);

  return 0;
}
