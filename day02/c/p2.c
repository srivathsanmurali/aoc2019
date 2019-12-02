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

int* modify_memory(int* memory, int noun, int verb) {
  int* new_memory = malloc(sizeof(int) * 128);
  for(int address = 0; address < 128; address++)
    new_memory[address] = memory[address];

  new_memory[1] = noun;
  new_memory[2] = verb;
  return new_memory;
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
  for(int noun = 0; noun < 100; noun++) {
    for(int verb = 0; verb < 100; verb++) {
      int* mm = modify_memory(memory, noun, verb);
      int result = run(mm);
      free(mm);
      if(result == 19690720) {
        printf("Result = %d\n", (100*noun)+verb);
        goto done;
      }
    }
  }
done:
  free(memory);

  return 0;
}
