#include <stdio.h>

#define numItems 3

int main()
{
	double ar[numItems];
	int i;

	for (i = 0; i < numItems; ++i)
		printf("Value = %f\n", ar[i]);

	return 0;
}
