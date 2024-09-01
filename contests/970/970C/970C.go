package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CF970C(l int, r int) int {

	increment := 1
	numIntegers := 1
	for l < r {
		l += increment
		if l <= r {
			numIntegers++
		}
		increment += 1
	}
	return numIntegers
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
		parts := strings.Split(line, " ")

		// Convert the split strings to integers
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])

		// Output or process the integers as needed
		writer.WriteString(fmt.Sprintf("%d\n", CF970C(a, b)))
	}
}
