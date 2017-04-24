// Utility to reverse a single line of text
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read line: %s", err)
		os.Exit(1)
	}
	i := 0
	j := len(line) - 1 - 1 // Extra -1 to exclude delimiter
	for i < j {
		t := line[i]
		line[i] = line[j]
		line[j] = t
		i++
		j--
	}
	fmt.Printf("%s", line)
}
