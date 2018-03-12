#include "grains.h"

long long square(int num)
{
   if ( num < 1 || num > 64) {
      return 0;
   }

   return (long long) 1 << (num - 1);
}

long long total(void)
{
   long long result = 0;

   for (long long i = 1; i <= 64; i++) {
      result += square(i);
   }

   return result;
}
