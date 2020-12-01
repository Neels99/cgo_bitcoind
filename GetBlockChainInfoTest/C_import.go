package main

/*
#include "test.h"

int TestFunc2(ts a){
	return a.a;
}
*/
import "C"
import "fmt"

func main() {
	fmt.Println("test")
	//C.TestFunc2()
	fmt.Println(int(C.TestFunc2(C.struct_TestStruct{5})))
}
