// show arguments of command line

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "skip symbol of new line, boolean")
var sep = flag.String("s", " ", "separator, string")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
