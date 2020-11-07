package main

//#include <jni.h>
import "C"
import (
	"fmt"
)

//export Java_VierBitAddierer_booleanTestMethod
func Java_VierBitAddierer_booleanTestMethod(env *C.JNIEnv, clazz C.jclass, a C.jboolean) {
	fmt.Printf("%v (%T)\n", a, a)
}

//export Java_VierBitAddierer_nand
func Java_VierBitAddierer_nand(env *C.JNIEnv, clazz C.jclass, a, b C.jboolean) C.jboolean{
	return nand(a,b)
}

//export Java_VierBitAddierer_xor
func Java_VierBitAddierer_xor(env *C.JNIEnv, clazz C.jclass, a, b C.jboolean) C.jboolean {
	return xor(a, b)
}

func nand(a, b C.jboolean) C.jboolean{
	if a & b > 0 {
		return 0
	}
	return 1
}

func xor(a, b C.jboolean) C.jboolean {
	c := nand(a,b)
	d, e := nand(a,c), nand(b,c)
	f := nand(d,e)
	return f
}

func main() {}
