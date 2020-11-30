#include <stdio.h>

typedef struct TestStruct{
    int a;
} ts;


int TS(ts a){
    return a.a;
}

void TestFunc(){
    printf("TESTFUNC C++");
}