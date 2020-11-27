#include <iostream>
#include "test.h"
using namespace std;

struct argss{
	int a;
	int b;
};


int main(){
	cout << "test" << endl;
	Start();
	int res = TestSUM(123,321);
	cout << res << endl;
	ARG argss2 = {1, 10};
	RES result = Sum2(argss2);
	cout << "result2: " << result.result << endl;
}

