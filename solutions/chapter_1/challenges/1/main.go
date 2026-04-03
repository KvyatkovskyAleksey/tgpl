// echo command line names and file name
package main

import (
	"fmt"
	"os"
)

func main() {
	var str, sep string
	for _, chunk := range os.Args {
		str += sep + chunk
		sep = " "
	}
	fmt.Println(str)
}
