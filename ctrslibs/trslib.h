#ifndef _TRSLIB_H
#define _TRSLIB_H


#ifdef __cplusplus
extern "C" {
#endif // __cplusplus

#include <stdio.h>
#include <stdbool.h>
#include <stdint.h>

#define min(a,b) \
   ({ __typeof__ (a) _a = (a); \
       __typeof__ (b) _b = (b); \
     _a < _b ? _a : _b; })

#define max(a,b) \
   ({ __typeof__ (a) _a = (a); \
       __typeof__ (b) _b = (b); \
     _a > _b ? _a : _b; })

typedef struct myStruct1 {
	char a[10];
} st1_t;

typedef struct myStruct2 {
	char a[8];
	char b[2];
} st2_t;

extern void myPrintFunction(char *s);
extern void myPrintFunction2(void);

extern st1_t s1;
extern st2_t s2;

#ifdef __cplusplus
}
#endif // __cplusplus


#endif /* TRSLIB_H */
