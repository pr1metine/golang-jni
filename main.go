// DO NOT REMOVE ANY COMMENTS
package main

//#include <jni.h>
//static jobject allocAddiererAusgabe(JNIEnv * env, jboolean s, jboolean ue) {
//		jclass clazz = (*env)->FindClass(env, "VierBitAddierer$AddiererAusgabe");
//		jobject obj = (*env)->AllocObject(env, clazz);
//		jfieldID fieldS = (*env)->GetFieldID(env, clazz, "s", "Z");
//		jfieldID fieldUE = (*env)->GetFieldID(env, clazz, "ue", "Z");
//		(*env)->SetBooleanField(env, obj, fieldS, s);
//		(*env)->SetBooleanField(env, obj, fieldUE, ue);
//		return obj;
//}
//
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

//export Java_VierBitAddierer_ha
func Java_VierBitAddierer_ha(env *C.JNIEnv, clazz C.jclass, a, b C.jboolean) C.jobject {
	s, ue := ha(a, b)
	return C.allocAddiererAusgabe(env, s, ue)
}

//export Java_VierBitAddierer_fa
func Java_VierBitAddierer_fa(env *C.JNIEnv, clazz C.jclass, a, b, ue C.jboolean) C.jobject {
	s2, ue3 := fa(a, b, ue)
	return C.allocAddiererAusgabe(env, s2, ue3)
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

func ha(a, b C.jboolean) ( C.jboolean, C.jboolean ) {
	return xor(a,b), a & b
}

func fa(a, b, ue C.jboolean) (C.jboolean, C.jboolean) {
	s1, ue1 := ha(a, b)
	s2, ue2 := ha(s1, ue)
	ue3 := ue1 | ue2
	return s2, ue3
}

func main() {}
