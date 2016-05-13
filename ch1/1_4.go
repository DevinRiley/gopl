package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	lineFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, lineFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, lineFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			for _, filename := range lineFiles[line] {
				fmt.Printf("%s\t", filename)
			}
			fmt.Println("")
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, lineFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		lineFiles[line] = addUniqueFileName(lineFiles[line], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}

func addUniqueFileName(files []string, filename string) []string {
	for _, file := range files {
		if file == filename {
			return files
		}
	}

	return append(files, filename)
}

//!-
