package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CF971B(beats []string) []string {
	notes := []string{}
	for i := len(beats) - 1; i >= 0; i-- {
		for j := 1; j <= len(beats[i]); j++ {
			if beats[i][j-1] == '#' {
				notes = append(notes, fmt.Sprint(j))
				break
			}
		}
	}
	return notes
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// Read the number of lines (test cases)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	numLines, _ := strconv.Atoi(line)

	// Process each line of input
	for i := 0; i < numLines; i++ {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		numBeats, _ := strconv.Atoi(line)
		beats := []string{}
		for i := 0; i < numBeats; i++ {
			line, _ = reader.ReadString('\n')
			line = strings.TrimSpace(line)
			beats = append(beats, line)
		}

		// Convert the split strings to integers

		// Output or process the integers as needed
		writer.WriteString(strings.Join(CF971B(beats), " ") + "\n")
	}
}
