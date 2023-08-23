package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func main() {
	// These 4 lines simply collect profile data into default.pgo file
	f, _ := os.Create("default.pgo")
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i <= 1_000_000; i++ {
		Func1()
	}
}

func Func1() {
	fmt.Println("print in func #1")
	Func2()
}

func Func2() {
	fmt.Println("print in func #2")
}
