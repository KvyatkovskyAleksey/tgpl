// show index and args on each string
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for idx, arg := range os.Args[1:] {
		fmt.Println(strconv.Itoa(idx) + " " + arg)
	}

}
