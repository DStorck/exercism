#include "difference_of_squares.h"
#include <stdio.h>

int sum_of_squares(int num)
{
  int total = 0 ;
  for ( int i = 1; i <= num; i++) {
    total +=  i * i;
  }
  return total;
}

int square_of_sum(int num)
{
  int total = 0 ;
  for ( int i = 1; i <= num; i++){
    total += i;
  }
  return total * total ;
}

int difference_of_squares(int num)
{
  return square_of_sum(num) - sum_of_squares(num);
}
