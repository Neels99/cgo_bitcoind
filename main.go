// package name: test
package main

/*
struct ARG {
    int a;
    int b;
};

struct RES {
    int result;
};
*/
import "C"

//export Sum2
func Sum2(arg C.struct_ARG) C.struct_RES {
	return C.struct_RES{arg.a + arg.b}
}

func main() {
	//println(Sum2(123, 3222))
}
