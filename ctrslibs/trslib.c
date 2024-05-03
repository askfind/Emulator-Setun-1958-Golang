#include <stdio.h>
#include <string.h>
#include "trslib.h"

void myPrintFunction(char *s) {
	printf("%s\n", s);
}

char dest[] = "oldstring";
const char src[] = "newstring";

st1_t s1 = {"0987654321"};
st2_t s2 = {"12345678","90"};

void myPrintFunction2(void) {
	printf("Hello from inline C\n");
	printf("%s\n",dest);
}

#if 0
/*
int main (void) {

	printf("Before memmove dest = %s, src = %s\n", dest, src);
	memmove(dest, src, 9);
	printf("After memmove dest = %s, src = %s\n", dest, src);

	st1_t s1 = {"0987654321"};
	st2_t s2 = {"12345678","90"};

	memset(&s2,0,sizeof(s2));
	memmove(&s2,&s1,sizeof(s1));

	//s2 = s1; //error
	printf(" - s1 %s\n",s1.a);
	printf(" - s2.a %s\n",s2.a);
	printf(" - s2.b %s\n",s2.b);
	printf(" - s2 %p\n",&s2);

	return 0;

}
*/
#endif
