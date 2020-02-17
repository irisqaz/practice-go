package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func func0() {}

func func2(a int, b string) {}

func funcOut(a int, b string) (int, error) { return 0, nil }

func funcVariadic(a int, b ...string) (int, error) { return 0, nil }

func main() {
	for _, f := range []interface{}{func0, func2, funcOut, funcVariadic} {
		name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		fmt.Println("name:", name)
		t := reflect.TypeOf(f)
		fmt.Println("type:", t, "in-params:", t.NumIn(), "out-params:", t.NumOut())
		in := make([]reflect.Type, t.NumIn())
		for i := range in {
			in[i] = t.In(i)
		}
		fmt.Println("inputs:", in)
		out := make([]reflect.Type, t.NumOut())
		for i := range out {
			out[i] = t.Out(i)
		}
		fmt.Println("outputs:", out)
		fmt.Println("variadic", t.IsVariadic())
		fmt.Println()
		// fmt.Printf("%q %v %v %v\n", name, in, out, t.IsVariadic())
	}
}
