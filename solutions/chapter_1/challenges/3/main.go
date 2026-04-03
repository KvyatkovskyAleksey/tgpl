// compare times using join and loop for string concatention
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	slice := generateSliceOfStrings(10000)
	fmt.Println("Time with iterator:")
	showTimeExecution(func() { joinSliceWithIterator(slice) })()
	fmt.Println("Time with join function:")
	showTimeExecution(func() { joinSliceWithJoin(slice) })()
}

func joinSliceWithJoin(slice []string) {
	strings.Join(slice, " - ")
}

func joinSliceWithIterator(slice []string) {
	res, sep := "", ""
	for _, str := range slice {
		res += sep + str
		sep = " "
	}
}

func generateSliceOfStrings(len int) []string {
	slice := make([]string, len)
	for i := range slice {
		slice[i] = "str"
	}
	return slice
}

func showTimeExecution(fn func()) func() {
	return func() {
		start := time.Now()
		fn()
		duration := time.Since(start)
		fmt.Printf("Execution time: %s\n", duration)
	}
}
