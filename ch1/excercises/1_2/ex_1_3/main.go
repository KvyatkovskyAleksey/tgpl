//function for compare execution time of version with concatenations and joins
// good to run it with many arg, for example with `go run main.go $(for i in $(seq 1 1000); do echo "arg$i"; done)`

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func func1() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func func2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	startFunc1 := time.Now()
	func1()
	firstFunctionTime := time.Since(startFunc1)
	startFunc2 := time.Now()
	func2()
	secondFunctionTime := time.Since(startFunc2)
	output1 := fmt.Sprintf("Time for 1 %d", firstFunctionTime)
	output2 := fmt.Sprintf("Time for 2 %d", secondFunctionTime)
	fmt.Println(output1)
	fmt.Println(output2)

}
