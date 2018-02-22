#include "hamming.h"
#include <stdio.h>
#include <string.h>


int compute(char *first, char *second){
  int diff = -1 ;
  if (first && second && (strlen(first) == strlen(second))) {
    diff = 0 ;
    for (unsigned int i = 0 ; i < strlen(first); i++) {
      if (first[i] != second[i] ) {
        diff++;
      }
    }
  }
  return diff;
}
