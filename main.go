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

//export Start
func Start() {
    println("Second Test")
}

type TestType struct {
	a int
	b int
}

//export TestSUM
func TestSUM(a int, b int) int{
	var t TestType
	t = TestType{a, b}
	return (t.a + t.b) / 2
}

//export Sum2
func Sum2(arg C.struct_ARG) C.struct_RES {
	return C.struct_RES{arg.a + arg.b}
}

//func test1(a C.int, b C.int) C.struct_ARG {
//	return C.struct_ARG{a, b}
//}

func main() {
	println(TestSUM(123, 3222))
}
