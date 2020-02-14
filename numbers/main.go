package main

import (
	"fmt"

	"github.com/irisqaz/practice-go/test"
)

func main() {

	fmt.Println()
	fmt.Println(test.Rjustified(42, test.Signature(Integer)))
	anInt := Integer()
	printResult(anInt)

	fmt.Println(test.Rjustified(42, test.Signature(Integer8)))
	anInt8 := Integer8()
	printResult(anInt8)

	fmt.Println(test.Rjustified(42, test.Signature(Integer16)))
	anInt16 := Integer16()
	printResult(anInt16)

	fmt.Println(test.Rjustified(42, test.Signature(Integer32)))
	anInt32 := Integer32()
	printResult(anInt32)

	fmt.Println(test.Rjustified(42, test.Signature(Integer64)))
	anInt64 := Integer64()
	printResult(anInt64)

	fmt.Println(test.Rjustified(42, test.Signature(UnsignedInt8)))
	anUint8 := UnsignedInt8()
	printResult(anUint8)

	fmt.Println(test.Rjustified(42, test.Signature(UnsignedInt64)))
	anUint64 := UnsignedInt64()
	printResult(anUint64)

	fmt.Println(test.Rjustified(42, test.Signature(Float32)))
	aFloat32 := Float32()
	printResult(aFloat32)

	fmt.Println(test.Rjustified(42, test.Signature(Float64)))
	aFloat64 := Float64()
	printResult(aFloat64)

	fmt.Println()
}
func printResult(got interface{}) {
	strIn := fmt.Sprintf("%25s", "()")
	strGot := test.Ljustified(10, got)
	strGot = test.Green(strGot)

	fmt.Print(strIn, "  ->  ", strGot)
	fmt.Println()
}
