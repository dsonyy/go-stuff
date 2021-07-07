package main

//#include <math.h>
//double number() {
//  return 1.5;
//}
import "C"

func main() {
	num1 := C.number()
	num2 := C.double(5)
	println("The number is", C.fmax(num1, num2))
}
