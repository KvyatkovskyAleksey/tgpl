//shows substrings presented in text more than once and shows number of occurrences

package main

import (
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for text, count := range counts {
		if count > 1 {
			println(text, count)
		}
	}
}
