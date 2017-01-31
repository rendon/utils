// Utility to reverse lines in a file
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var lines [][]byte
	for {
		// Ignoring isPrefix, which should only affect for really long lines
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		lines = append(lines, line)
	}
	for i := len(lines) - 1; i >= 0; i-- {
		fmt.Printf("%s\n", lines[i])
	}
}
