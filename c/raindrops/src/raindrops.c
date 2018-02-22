#include "raindrops.h"
#include "stdio.h"
#include "string.h"

char *convert(char rain[], int drops)
{
   if (drops % 3 == 0) {
      strcat(rain, "Pling");
   }
   if (drops % 5 == 0) {
      strcat(rain, "Plang");
   }
   if (drops % 7 == 0) {
      strcat(rain, "Plong");
   }
   if (strlen(rain) == 0 ){
     sprintf(rain, "%d", drops);
   }
   return rain;
}
