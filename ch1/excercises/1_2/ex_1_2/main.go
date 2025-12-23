// print command line arguments with index and each on new string
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	sep = "\n"
	for index, arg := range os.Args[1:] {
		argString := fmt.Sprintf("%d: %s", index, arg)
		s += argString + sep
	}
	fmt.Println(s)
}
