// using dup2 as base show also which files have duplicates
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	counts := make(map[string]int)
	linesByFiles := make(map[string]map[string]struct{})
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, linesByFiles)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			countLines(f, counts, linesByFiles)
			f.Close()
		}
	}
	for line, count := range counts {
		if count > 1 {
			files_output, sep := "", ""
			for filename := range linesByFiles[line] {
				files_output += sep + filename
				sep = ", "
			}
			fmt.Printf("Line: \"%v\": %d, presented in files: %s\n", line, count, files_output)

		}
	}

}

func countLines(f *os.File, counts map[string]int, linesByFiles map[string]map[string]struct{}) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		counts[text] += 1
		if _, exists := linesByFiles[text]; !exists {
			linesByFiles[text] = map[string]struct{}{}
		}
		linesByFiles[text][filepath.Base(f.Name())] = struct{}{}
	}
}
